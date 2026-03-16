package scene

import (
	"ecs-test/client/config"
	broadcastService "ecs-test/client/middleware/broadcast/service"
	"ecs-test/client/middleware/socket/service"
	"ecs-test/client/system"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/events"
)

type WorldDisplaySceneFactory struct {
	w                donburi.World
	config           config.GameClientConfig
	socketService    *service.SocketService
	broadcastService *broadcastService.BroadcastService
	switcher         *SceneSwitcher
}

func NewWorldDisplaySceneFactory(
	w donburi.World,
	config config.GameClientConfig,
	socketService *service.SocketService,
	broadcastService *broadcastService.BroadcastService,
) *WorldDisplaySceneFactory {
	return &WorldDisplaySceneFactory{
		w:                w,
		config:           config,
		socketService:    socketService,
		broadcastService: broadcastService,
	}
}

func (w *WorldDisplaySceneFactory) Load() Scene {
	return NewWorldDisplayScene(
		w.w,
		w.config,
		w.socketService,
		w.broadcastService,
	)
}

type WorldDisplayScene struct {
	ecs *ecs.ECS
}

func NewWorldDisplayScene(
	w donburi.World,
	config config.GameClientConfig,
	socketService *service.SocketService,
	broadcastService *broadcastService.BroadcastService,
) *WorldDisplayScene {
	// systems
	width, height := config.Width, config.Height
	debugSystem := system.NewDebugSystem()
	worldMapLoader := system.NewWorldMapLoader(width, height)
	worldMapRender := system.NewWorldMapRender(width, height)
	cameraSystem := system.NewCamera(width, height, 2.0, 3.0)
	cameraBoundary := system.NewCameraBoundary(width, height)
	cameraFollow := system.NewCameraFollow(width, height, 0.2)
	player := system.NewPlayerSystem()
	playerDust := system.NewPlayerDustTrace()
	pathNavigation := system.NewPathNavigation()
	cursor := system.NewCursorSystem()

	movementBroadcast := system.NewMovementBroadcast(broadcastService)
	remotePlayers := system.NewRemotePlayers(socketService, broadcastService)
	connectionMonitor := system.NewConnectionMonitor(socketService, broadcastService, w, width, height)
	//ui := system.NewGameUi(width, height)

	layerIndex := &LayerIndex{}
	worldDisplayScene := &WorldDisplayScene{

		ecs: ecs.NewECS(w).
			AddSystem(cursor.Update).
			//AddSystem(ui.Update).
			AddSystem(debugSystem.Update).
			AddSystem(worldMapRender.Update).
			AddSystem(worldMapLoader.Update).
			AddSystem(cameraFollow.Update).
			AddSystem(cameraBoundary.Update).
			AddSystem(cameraSystem.Update).
			AddSystem(cameraBoundary.Update).
			AddSystem(player.Update).
			AddSystem(playerDust.Update).
			AddSystem(pathNavigation.Update).
			AddSystem(connectionMonitor.Update).
			AddSystem(movementBroadcast.Update).
			AddSystem(remotePlayers.Update).
			AddRenderer(layerIndex.Next(), worldMapLoader.Draw).
			AddRenderer(layerIndex.Next(), worldMapRender.DrawGround).
			AddRenderer(layerIndex.Next(), worldMapRender.DrawElevation).
			AddRenderer(layerIndex.Next(), worldMapRender.DrawDecorations).
			AddRenderer(layerIndex.Next(), pathNavigation.Draw).
			AddRenderer(layerIndex.Next(), playerDust.Draw).
			AddRenderer(layerIndex.Next(), player.Draw).
			AddRenderer(layerIndex.Next(), remotePlayers.Draw).
			AddRenderer(layerIndex.Next(), worldMapRender.DrawOtherDecorations).
			AddRenderer(layerIndex.Next(), worldMapRender.DrawWalkable).
			AddRenderer(layerIndex.Next(), worldMapRender.DrawShutter).
			AddRenderer(layerIndex.Next(), debugSystem.Draw).
			AddRenderer(layerIndex.Next(), connectionMonitor.Draw).
			//AddRenderer(layerIndex.Next(), cameraSystem.Draw).
			AddRenderer(layerIndex.Next(), cursor.Draw),
	}

	return worldDisplayScene
}

func (w *WorldDisplayScene) Update() error {
	w.ecs.Update()
	events.ProcessAllEvents(w.ecs.World)

	return nil
}

func (w *WorldDisplayScene) Draw(screen *ebiten.Image) {
	w.ecs.Draw(screen)
}
