package system

import (
	"ecs-test/client/component"
	"ecs-test/client/event"
	broadcastService "ecs-test/client/middleware/broadcast/service"
	socketService "ecs-test/client/middleware/socket/service"
	"ecs-test/client/view"
	"ecs-test/client/view/effects"
	"ecs-test/client/view/widgets"
	"ecs-test/shared/socket/payload"
	"github.com/charmbracelet/log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/furex/v2"
	"time"
)

type ConnectionMonitor struct {
	socketService    *socketService.SocketService
	broadcastService *broadcastService.BroadcastService
	world            donburi.World

	sessionEntry         *donburi.Entry
	playerLock           *donburi.Entry
	fadeToGray           *effects.FadeToGray
	shutter              *effects.Shutter
	shutterRunning       bool
	rootView             *furex.View
	connectionLostDialog *view.ConnectionLostDialog

	connected    bool
	progressView *furex.View
	progressBar  *widgets.ProgressBar
}

func NewConnectionMonitor(
	socketService *socketService.SocketService,
	broadcastService *broadcastService.BroadcastService,
	world donburi.World,
	screenWidth, screenHeight int,
) *ConnectionMonitor {

	cm := &ConnectionMonitor{
		socketService:    socketService,
		broadcastService: broadcastService,
		world:            world,

		shutter:    effects.NewShutter(screenWidth, screenHeight, 200, 1),
		fadeToGray: effects.NewFadeToGray(screenWidth, screenHeight),
		connected:  true,
		rootView: &furex.View{
			Width:        screenWidth,
			Height:       screenHeight,
			Direction:    furex.Column,
			AlignContent: furex.AlignContentCenter,
			Justify:      furex.JustifyCenter,
			AlignItems:   furex.AlignItemCenter,
		},
	}
	socketService.OnConnectionClosed(cm.connectionClosed)
	return cm.initView()
}

func (m *ConnectionMonitor) initView() *ConnectionMonitor {
	m.connectionLostDialog = view.NewConnectionLostDialog(
		2.0,
		m.reconnect,
		m.back,
	)
	m.rootView.AddChild(m.connectionLostDialog.View())

	// progress bar
	m.progressView = &furex.View{
		Direction:    furex.Row,
		WidthInPct:   100,
		AlignContent: furex.AlignContentCenter,
		Justify:      furex.JustifyCenter,
		AlignItems:   furex.AlignItemCenter,
		MarginTop:    32,
		Hidden:       true,
	}
	m.progressBar = widgets.NewProgressBar(0)
	m.progressView.AddChild(m.progressBar.View())

	m.rootView.AddChild(m.progressView)

	return m
}

func (m *ConnectionMonitor) connectionClosed(err error) {
	m.connected = false
	m.connectionLostDialog.Show()
	m.fadeToGray.FadeToGray()

	m.broadcastService.Disconnect()
	m.lockPlayer()
}

func (m *ConnectionMonitor) Update(ecs *ecs.ECS) {
	m.findSession(ecs)
	m.findPlayerLock(ecs)

	if m.shutterRunning {
		m.shutter.Update()
	}

	if m.connected {
		return
	}

	m.rootView.Update()
	m.fadeToGray.Update()
}

func (m *ConnectionMonitor) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	if m.connected {
		return
	}

	m.fadeToGray.Draw(screen, screen)
	m.rootView.Draw(screen)
	m.shutter.Draw(screen)
}

func (m *ConnectionMonitor) reconnect() {
	m.updateConnectionProgress(0)
	m.showProgress()

	err := m.socketService.Connect()
	if err != nil {
		m.hideProgress()
		log.Error(err)
		return
	}

	err = m.broadcastService.Connect()
	if err != nil {
		m.hideProgress()
		log.Error(err)
		return
	}

	session := component.Session.Get(m.sessionEntry)
	joinResultChan, err := m.socketService.RequestJoin(session.SessionId())

	if err != nil {
		m.hideProgress()
		return
	}

	go m.waitForJoinResult(joinResultChan)
}

func (m *ConnectionMonitor) back() {
	m.shutterRunning = true
	m.shutter.Start(effects.ShutterShrink, func() {
		m.shutterRunning = false
		event.SceneChangeEvent.Publish(
			m.world,
			event.NewSceneChange("menu"),
		)

		m.fadeToGray.FadeToNormal()
		m.connectionLostDialog.Hide()
		m.unlockPlayer()
	})
}

func (m *ConnectionMonitor) findSession(ecs *ecs.ECS) {
	if m.sessionEntry == nil {
		entry, ok := component.SessionQuery.First(ecs.World)
		if !ok {
			// when not found
		}
		m.sessionEntry = entry
	}
}

func (m *ConnectionMonitor) findPlayerLock(e *ecs.ECS) {
	if m.playerLock == nil {
		lockEntry, found := component.PlayerLockQuery.First(e.World)
		if !found {

		}
		m.playerLock = lockEntry
	}
}

func (m *ConnectionMonitor) waitForJoinResult(resultChan chan *payload.JoinResponse) {
	t := time.NewTicker(time.Second)
	defer t.Stop()

	maxSecondsToWait := 6
	secondsWaiting := 0
	timeoutChan := time.After(6 * time.Second) // Create the timeout channel outside the loop

	for {
		select {
		case joinResult := <-resultChan:
			m.progressBar.UpdateProgress(100)
			m.hideProgress()

			if joinResult.Status != payload.JoinSuccessful {
				err := m.socketService.Disconnect()
				if err != nil {
					log.Error(err)
				}
				m.broadcastService.Disconnect()
				return
			}

			m.fadeToGray.FadeToNormal()
			m.connected = true
			m.broadcastService.Start()
			m.unlockPlayer()
			return
		case <-t.C:
			secondsWaiting++
			percentage := int((1.0 - float64(maxSecondsToWait-secondsWaiting)/float64(maxSecondsToWait)) * 100)
			m.updateConnectionProgress(percentage)
		case <-timeoutChan:
			m.updateConnectionProgress(0)
			m.hideProgress()
			return
		}
	}
}

func (m *ConnectionMonitor) updateConnectionProgress(progress int) {
	m.progressBar.UpdateProgress(progress)
}

func (m *ConnectionMonitor) showProgress() {
	m.progressView.Hidden = false
	m.connectionLostDialog.DisableRetryButton()
}

func (m *ConnectionMonitor) hideProgress() {
	m.progressView.Hidden = true
	m.connectionLostDialog.EnableRetryButton()
}

func (m *ConnectionMonitor) lockPlayer() {
	if m.playerLock == nil {
		log.Error("Unable to lock!")
		return
	}

	lock := component.PlayerLock.Get(m.playerLock)
	lock.Lock()
}

func (m *ConnectionMonitor) unlockPlayer() {
	if m.playerLock == nil {
		log.Error("Unable to lock!")
		return
	}

	lock := component.PlayerLock.Get(m.playerLock)
	lock.Unlock()
}
