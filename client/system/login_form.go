package system

import (
	"ecs-test/client/component"
	broadcastService "ecs-test/client/middleware/broadcast/service"
	playerService "ecs-test/client/middleware/player/service"
	"ecs-test/client/middleware/session/service"
	socketService "ecs-test/client/middleware/socket/service"
	"ecs-test/client/view"
	"ecs-test/shared/socket/payload"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"github.com/yohamta/furex/v2"
	"time"
)

type LoginFormSystem struct {
	sessionService   *service.SessionService
	socketService    *socketService.SocketService
	broadcastService *broadcastService.BroadcastService
	playerService    *playerService.PlayerService

	loginForm   *view.LoginForm
	errorDialog *view.ErrorDialog
	ui          *furex.View

	sessionQuery *donburi.Query
	sessionEntry *donburi.Entry

	currentlyLoggingIn bool
	visible            bool

	onClose    func()
	afterLogin func()
}

func NewLoginFormSystem(
	sessionService *service.SessionService,
	socketService *socketService.SocketService,
	broadcastService *broadcastService.BroadcastService,
	playerService *playerService.PlayerService,
	width, height int,
) *LoginFormSystem {

	sys := &LoginFormSystem{
		ui: &furex.View{
			Width:        width,
			Height:       height,
			Direction:    furex.Column,
			AlignContent: furex.AlignContentCenter,
			Justify:      furex.JustifyCenter,
			AlignItems:   furex.AlignItemCenter,
		},

		sessionService:   sessionService,
		socketService:    socketService,
		broadcastService: broadcastService,
		playerService:    playerService,

		currentlyLoggingIn: false,
		visible:            false,

		sessionQuery: query.NewQuery(filter.Contains(component.Session)),
	}
	return sys.initView()
}

func (l *LoginFormSystem) initView() *LoginFormSystem {
	l.loginForm = view.NewLoginForm(2.0, l.login)
	l.ui.AddChild(l.loginForm.View())

	l.errorDialog = view.NewErrorDialog(2.0)
	l.ui.AddChild(l.errorDialog.View())

	return l
}

func (l *LoginFormSystem) Update(ecs *ecs.ECS) {
	l.findSessionComponent(ecs)

	if !l.visible || l.currentlyLoggingIn {
		return
	}

	l.ui.Update()
	l.loginForm.UpdateShake()

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		l.Hide()
		l.Reset()

		if l.onClose != nil {
			l.onClose()
		}
	}
}

func (l *LoginFormSystem) findSessionComponent(ecs *ecs.ECS) {
	if l.sessionEntry == nil {
		entry, ok := l.sessionQuery.First(ecs.World)
		if ok {
			l.sessionEntry = entry
		}
	}
}

func (l *LoginFormSystem) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	if !l.visible {
		return
	}

	l.ui.Draw(screen)
}

func (l *LoginFormSystem) Show(onClose func(), afterLogin func()) {
	l.visible = true
	l.onClose = onClose
	l.afterLogin = afterLogin
}

func (l *LoginFormSystem) login(email, password string) {
	l.currentlyLoggingIn = true
	l.errorDialog.Hide()
	l.loginForm.ShowProgress()

	err := l.socketService.Connect()
	if err != nil {
		log.Error(err)
		l.cancelLogin()
		return
	}

	err = l.broadcastService.Connect()
	if err != nil {
		log.Error(err)
		l.cancelLogin()
		return
	}

	sessionChan := l.sessionService.CreateSessionForCredentials(email, password)
	go l.waitForSession(sessionChan)
}

func (l *LoginFormSystem) waitForSession(sessionChan chan service.CreateSessionResult) {
	t := time.NewTicker(time.Second)
	defer t.Stop()

	maxSecondsToWait := 6
	secondsWaiting := 0
	timeoutChan := time.After(6 * time.Second) // Create the timeout channel outside the loop
	for {
		select {
		case createSessionResult := <-sessionChan:
			if createSessionResult.Failed() {
				l.handleSessionFailure(createSessionResult)
				return
			}

			// After successfully acquiring the session ID
			session := component.Session.Get(l.sessionEntry)
			session.SetSessionId(createSessionResult.SessionId)

			joinResultChan, err := l.socketService.RequestJoin(createSessionResult.SessionId)
			if err != nil {
				l.cancelLogin()
				return
			}
			l.waitForJoinResult(joinResultChan, &secondsWaiting, maxSecondsToWait, t)

			return

		case <-t.C:
			l.updateProgress(&secondsWaiting, maxSecondsToWait)
		case <-timeoutChan:
			l.cancelLogin()
			return
		}
	}
}

// Handles when the session creation fails
func (l *LoginFormSystem) handleSessionFailure(result service.CreateSessionResult) {
	l.loginForm.HideProgress()
	l.currentlyLoggingIn = false

	switch result.Status {
	case service.CreateSessionUnauthorized:
		l.loginForm.Shake(200*time.Millisecond, 4)
	default:
		l.errorDialog.Show()
	}
}

// Waits for the join result after successfully acquiring the session ID
func (l *LoginFormSystem) waitForJoinResult(
	joinResultChan chan *payload.JoinResponse,
	secondsWaiting *int,
	maxSecondsToWait int,
	t *time.Ticker,
) {
	for {
		select {
		case joinResult := <-joinResultChan:
			if joinResult.Status != payload.JoinSuccessful {
				l.loginForm.HideProgress()
				l.currentlyLoggingIn = false
				l.loginForm.Shake(200*time.Millisecond, 4)
				return
			}

			fmt.Printf("Join successful, server time: %s\n", joinResult.Time.String())
			if len(joinResult.OtherPlayers) > 0 {
				for _, otherPlayer := range joinResult.OtherPlayers {
					l.playerService.SpawnRemotePlayer(otherPlayer.SessionId, otherPlayer.Position.X, otherPlayer.Position.Y)
				}
			}

			l.loginForm.UpdateProgress(100)
			l.playerService.SpawnPlayer(joinResult.Position.X, joinResult.Position.Y)
			l.broadcastService.Start()

			if l.afterLogin != nil {
				go l.afterLogin()
			}
			return
		case <-t.C:
			l.updateProgress(secondsWaiting, maxSecondsToWait)
		}
	}
}

func (l *LoginFormSystem) updateProgress(secondsWaiting *int, maxSecondsToWait int) {
	*secondsWaiting++
	percentage := int((1.0 - float64(maxSecondsToWait-*secondsWaiting)/float64(maxSecondsToWait)) * 100)
	l.loginForm.UpdateProgress(percentage)
}

func (l *LoginFormSystem) cancelLogin() {
	l.errorDialog.Show()

	l.loginForm.HideProgress()
	l.currentlyLoggingIn = false
}

func (l *LoginFormSystem) Hide() {
	l.visible = false
	l.currentlyLoggingIn = false
}

func (l *LoginFormSystem) Reset() {
	l.currentlyLoggingIn = false
	l.loginForm.HideProgress()
	l.loginForm.Reset()
	l.errorDialog.Hide()
}
