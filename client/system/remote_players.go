package system

import (
	"ecs-test/client/archetype"
	"ecs-test/client/component"
	broadcastService "ecs-test/client/middleware/broadcast/service"
	"ecs-test/client/middleware/socket/service"
	"ecs-test/shared/socket/payload"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/ganim8/v2"
)

type RemotePlayers struct {
	initialized   bool
	world         donburi.World
	socketService *service.SocketService

	camera           *donburi.Entry
	readingMovements bool
	movementChan     chan *payload.PlayerMovementBroadcast
	broadcastService *broadcastService.BroadcastService
}

func NewRemotePlayers(socketService *service.SocketService, broadcastService *broadcastService.BroadcastService) *RemotePlayers {
	return &RemotePlayers{
		socketService:    socketService,
		broadcastService: broadcastService,
		movementChan:     make(chan *payload.PlayerMovementBroadcast, 20),
	}
}

func (p *RemotePlayers) Update(e *ecs.ECS) {
	p.findCamera(e)

	if !p.initialized {
		p.world = e.World
		p.socketService.OnPlayerJoined(p.playerJoined)
		p.broadcastService.OnPlayerMovement(p.playerMoved)
		p.initialized = true
	}

	component.RemotePlayerQuery.Each(e.World, func(entry *donburi.Entry) {
		remotePlayer := component.RemotePlayer.Get(entry)
		playerTransform := transform.Transform.Get(entry)
		remotePlayer.Update(playerTransform, e.Time.DeltaTime().Seconds())
	})

	// process buffered movements
	for i := 0; i < 100; i++ {
		select {
		case movement := <-p.movementChan:
			p.processMovement(e.World, movement)
		default:
			// Exit early if no more movements in the channel
			return
		}
	}
}

func (p *RemotePlayers) processMovement(world donburi.World, movement *payload.PlayerMovementBroadcast) {
	component.RemotePlayerQuery.Each(world, func(entry *donburi.Entry) {
		remotePlayer := component.RemotePlayer.Get(entry)
		if remotePlayer.SessionId != movement.SessionId {
			return
		}
		if remotePlayer.LatestUpdateAfter(movement.Time) {
			return
		}

		playerTransform := transform.Transform.Get(entry)

		dir := component.PlayerDirectionUp
		switch movement.Direction {
		case 1:
			dir = component.PlayerDirectionDown
		case 3:
			dir = component.PlayerDirectionRight
		case 2:
			dir = component.PlayerDirectionLeft
		}

		remotePlayer.Move(dir, playerTransform, movement.X, movement.Y)
		remotePlayer.UpdateMovementTime(movement.Time)
	})

}

func (p *RemotePlayers) findCamera(e *ecs.ECS) {
	if p.camera == nil {
		cameraEntity, cameraFound := component.CameraQuery.First(e.World)
		if !cameraFound {
		}
		p.camera = cameraEntity
	}
}

func (p *RemotePlayers) playerJoined(joinBroadcast *payload.JoinBroadcast) {
	log.Infof(
		"Player joined: %s, at location: %f, %f\n",
		joinBroadcast.UserName,
		joinBroadcast.Position.X,
		joinBroadcast.Position.Y,
	)

	position := joinBroadcast.Position
	_ = archetype.NewRemotePlayer(
		p.world,
		joinBroadcast.SessionId,
		math.NewVec2(position.X, position.Y),
	)
}

func (p *RemotePlayers) playerMoved(movementBroadcast *payload.PlayerMovementBroadcast) {
	fmt.Printf("Player moved: %s\n", movementBroadcast.SessionId)
	p.movementChan <- movementBroadcast
}

func (p *RemotePlayers) Draw(e *ecs.ECS, screen *ebiten.Image) {
	if p.camera == nil {
		return
	}
	cameraTransform := transform.Transform.Get(p.camera)

	component.RemotePlayerQuery.Each(e.World, func(entry *donburi.Entry) {
		playerData := component.RemotePlayer.Get(entry)
		playerTransform := transform.Transform.Get(entry)

		playerX, playerY := playerTransform.LocalPosition.
			Sub(cameraTransform.LocalPosition).
			Mul(cameraTransform.LocalScale).
			XY()

		drawOpts := ganim8.DrawOpts(playerX, playerY, 0, cameraTransform.LocalScale.X, cameraTransform.LocalScale.Y)
		playerData.DrawCurrentAnimation(screen, drawOpts)
	})
}
