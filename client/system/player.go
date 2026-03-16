package system

import (
	"ecs-test/client/component"
	"ecs-test/client/event"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/ganim8/v2"
	"image/color"
	"log"
)

type PlayerSystem struct {
	debugQuery *donburi.Query

	playerEntry *donburi.Entry
	debugEntry  *donburi.Entry
	worldMap    *donburi.Entry
	camera      *donburi.Entry
	playerLock  *donburi.Entry
}

func NewPlayerSystem() *PlayerSystem {
	return &PlayerSystem{
		debugQuery: donburi.NewQuery(filter.Contains(component.Debug)),
	}
}

func (ps *PlayerSystem) Update(e *ecs.ECS) {
	ps.findPlayer(e)
	ps.findPlayerLock(e)
	ps.findDebugComponent(e)
	ps.findCamera(e)
	ps.findWorldMap(e)

	if ps.camera == nil || ps.playerEntry == nil {
		return
	}

	worldMap := component.WorldMap.Get(ps.worldMap)
	if !worldMap.IsLoaded() {
		return
	}

	cameraData := component.Camera.Get(ps.camera)
	playerTransform := transform.Transform.Get(ps.playerEntry)
	playerData := component.Player.Get(ps.playerEntry)
	playerData.Update(playerTransform, worldMap.WalkableLayer.TileWidth, worldMap.WalkableLayer.TileHeight)

	if ps.playerLocked() {
		return
	}

	isMoving := false
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		playerData.MoveUp(playerTransform, ps.validCoordinates)
		cameraData.FollowPlayer()
		isMoving = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		playerData.MoveDown(playerTransform, ps.validCoordinates)
		cameraData.FollowPlayer()
		isMoving = true

	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		playerData.MoveLeft(playerTransform, ps.validCoordinates)
		cameraData.FollowPlayer()
		isMoving = true

	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		playerData.MoveRight(playerTransform, ps.validCoordinates)
		cameraData.FollowPlayer()
		isMoving = true
	}

	if isMoving || playerData.FollowingPath() {
		event.PlayerMovementEvent.Publish(e.World, event.NewPlayerMovement(
			playerData.LookDirection(),
			playerTransform.LocalPosition.X,
			playerTransform.LocalPosition.Y,
		))
	}

	if !isMoving && !playerData.FollowingPath() {
		playerData.GoIdle()
	}
}

func (ps *PlayerSystem) validCoordinates(x float64, y float64) bool {
	worldMap := component.WorldMap.Get(ps.worldMap)
	return worldMap.WalkableLayer.IsValidCoordinate(x, y)
}

func (ps *PlayerSystem) findCamera(e *ecs.ECS) {
	if ps.camera == nil {
		cameraEntity, cameraFound := component.CameraQuery.First(e.World)
		if !cameraFound {
			// Handle case where the world map entity is not found
		}
		ps.camera = cameraEntity
	}
}

func (ps *PlayerSystem) findDebugComponent(e *ecs.ECS) {
	if ps.debugEntry == nil {
		debugEntry, entryFound := ps.debugQuery.First(e.World)
		if !entryFound {
			log.Fatalf("Debug entry not found!")
			return
		}
		ps.debugEntry = debugEntry
	}
}

func (ps *PlayerSystem) findPlayer(e *ecs.ECS) {
	if ps.playerEntry == nil {
		playerEntry, entryFound := component.PlayerQuery.First(e.World)
		if !entryFound {
			//log.Fatalf("Player entry not found!")
			return
		}
		ps.playerEntry = playerEntry
	}
}

func (ps *PlayerSystem) findPlayerLock(e *ecs.ECS) {
	if ps.playerLock == nil {
		lockEntry, found := component.PlayerLockQuery.First(e.World)
		if !found {

		}
		ps.playerLock = lockEntry
	}
}

func (ps *PlayerSystem) findWorldMap(e *ecs.ECS) {
	if ps.worldMap == nil {
		entry, ok := component.WorldMapQuery.First(e.World)
		if ok {
			ps.worldMap = entry
		}
	}
}

func (ps *PlayerSystem) playerLocked() bool {
	if ps.playerLock == nil {
		return true
	}

	lock := component.PlayerLock.Get(ps.playerLock)
	return lock.IsLocked()
}

func (ps *PlayerSystem) Draw(e *ecs.ECS, screen *ebiten.Image) {
	if ps.camera == nil || ps.playerEntry == nil {
		return
	}

	worldMap := component.WorldMap.Get(ps.worldMap)
	if !worldMap.IsLoaded() {
		return
	}

	cameraTransform := transform.Transform.Get(ps.camera)

	playerData := component.Player.Get(ps.playerEntry)
	playerTransform := transform.Transform.Get(ps.playerEntry)

	playerX, playerY := playerTransform.LocalPosition.
		Sub(cameraTransform.LocalPosition).
		Mul(cameraTransform.LocalScale).
		XY()

	drawOpts := ganim8.DrawOpts(playerX, playerY, 0, cameraTransform.LocalScale.X, cameraTransform.LocalScale.Y)
	playerData.DrawCurrentAnimation(screen, drawOpts)

	debugData := component.Debug.Get(ps.debugEntry)
	if debugData.IsEnabled() {
		vector.DrawFilledCircle(screen, float32(playerX), float32(drawOpts.Y), 5, color.RGBA{R: 255}, false)
	}
}
