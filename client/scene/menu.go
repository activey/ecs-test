package scene

import (
	"ecs-test/client/config"
	"ecs-test/client/event"
	broadcastService "ecs-test/client/middleware/broadcast/service"
	playerService "ecs-test/client/middleware/player/service"
	"ecs-test/client/middleware/session/service"
	socketService "ecs-test/client/middleware/socket/service"
	"ecs-test/client/system"
	"ecs-test/client/view/effects"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/events"
)

type MenuSceneFactory struct {
	world            donburi.World
	config           config.GameClientConfig
	sessionService   *service.SessionService
	socketService    *socketService.SocketService
	broadcastService *broadcastService.BroadcastService
	playerService    *playerService.PlayerService
	sceneSwitcher    *SceneSwitcher
}

func NewMenuSceneFactory(
	world donburi.World,
	config config.GameClientConfig,
	sessionService *service.SessionService,
	socketService *socketService.SocketService,
	broadcastService *broadcastService.BroadcastService,
	playerService *playerService.PlayerService,
	sceneSwitcher *SceneSwitcher,
) *MenuSceneFactory {
	return &MenuSceneFactory{
		world:            world,
		config:           config,
		sessionService:   sessionService,
		socketService:    socketService,
		broadcastService: broadcastService,
		playerService:    playerService,
		sceneSwitcher:    sceneSwitcher,
	}
}

func (m *MenuSceneFactory) Load() Scene {
	return NewMenuScene(
		m.world,
		m.config,
		m.sessionService,
		m.socketService,
		m.broadcastService,
		m.playerService,
		m.sceneSwitcher,
	)
}

type MenuScene struct {
	ecs       *ecs.ECS
	shutter   *system.ShutterSystem
	loginForm *system.LoginFormSystem
	menu      *system.MenuWithBackground

	sceneSwitcher *SceneSwitcher

	uiLocked      bool
	quitRequested bool
}

func NewMenuScene(
	world donburi.World,
	config config.GameClientConfig,
	sessionService *service.SessionService,
	socketService *socketService.SocketService,
	broadcastService *broadcastService.BroadcastService,
	playerService *playerService.PlayerService,
	sceneSwitcher *SceneSwitcher,
) *MenuScene {
	// systems
	width, height := config.Width, config.Height
	debugSystem := system.NewDebugSystem()
	loginForm := system.NewLoginFormSystem(sessionService, socketService, broadcastService, playerService, width, height)
	shutterSystem := system.NewShutterSystem(width, height, 200, 1)
	cursor := system.NewCursorSystem()
	menu := system.NewMenuWithBackground(world, width, height)

	layerIndex := &LayerIndex{}
	menuScene := &MenuScene{
		shutter:   shutterSystem,
		loginForm: loginForm,
		menu:      menu,

		sceneSwitcher: sceneSwitcher,

		ecs: ecs.NewECS(world).
			AddSystem(cursor.Update).
			AddSystem(debugSystem.Update).
			AddSystem(menu.Update).
			AddSystem(loginForm.Update).
			AddSystem(shutterSystem.Update).
			AddRenderer(layerIndex.Next(), menu.Draw).
			AddRenderer(layerIndex.Next(), loginForm.Draw).
			AddRenderer(layerIndex.Next(), shutterSystem.Draw).
			AddRenderer(layerIndex.Next(), debugSystem.Draw).
			AddRenderer(layerIndex.Next(), cursor.Draw),
	}

	menuScene.shutter.Start(effects.ShutterExpand, func() {
		//w.shutterRunning = false
	})

	event.MenuSelectionEvent.Subscribe(world, menuScene.handleMenuSelection)
	return menuScene
}

func (m *MenuScene) Update() error {
	if m.quitRequested {
		return ebiten.Termination
	}

	events.ProcessAllEvents(m.ecs.World)
	m.ecs.Update()
	return nil
}

func (m *MenuScene) Draw(screen *ebiten.Image) {
	m.ecs.Draw(screen)
}

func (m *MenuScene) handleMenuSelection(w donburi.World, input event.MenuSelection) {
	if m.uiLocked {
		print("locked!")
		return
	}

	m.uiLocked = true

	switch input {
	case event.Quit:
		m.quitRequested = true
	case event.JoinWorld:
		m.joinWorld(w)
	}
}

func (m *MenuScene) joinWorld(w donburi.World) {
	m.menu.FadeToGray()

	m.loginForm.Show(m.loginFormClosed, func() {
		m.shutter.Start(effects.ShutterShrink, func() {
			m.sceneSwitcher.Switch("world_display")
			// for some reason there is some strange glitch when using events here
			//event.SceneChangeEvent.Publish(w, event.NewSceneChange("world_display"))
			m.disposeLoginForm()
		})
	})
}

func (m *MenuScene) loginFormClosed() {
	m.menu.FadeToNormal()
	m.uiLocked = false
}

func (m *MenuScene) disposeLoginForm() {
	m.loginForm.Reset()
	m.loginForm.Hide()
	m.menu.FadeToNormal()
	m.uiLocked = false
}
