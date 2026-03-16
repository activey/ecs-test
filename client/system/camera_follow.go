package system

import (
	"ecs-test/client/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
)

type CameraFollow struct {
	camera   *donburi.Entry
	worldMap *donburi.Entry
	player   *donburi.Entry

	screenWidth     float64
	screenHeight    float64
	smoothingFactor float64
}

func NewCameraFollow(
	screenWidth, screenHeight int,
	smoothingFactor float64,
) *CameraFollow {
	return &CameraFollow{
		screenWidth:     float64(screenWidth),
		screenHeight:    float64(screenHeight),
		smoothingFactor: smoothingFactor,
	}
}

func (f *CameraFollow) Update(e *ecs.ECS) {
	f.findPlayer(e)
	f.findCamera(e)
	f.findWorldMap(e)

	if f.camera == nil || f.worldMap == nil {
		return
	}

	cameraData := component.Camera.Get(f.camera)
	if !cameraData.ShouldFollowPlayer() {
		return
	}

	cameraTransform := transform.Transform.Get(f.camera)
	playerTransform := transform.Transform.Get(f.player)

	playerX, playerY := playerTransform.LocalPosition.XY()
	smoothingFactor := f.smoothingFactor
	targetX := playerX - (f.screenWidth/2.0)/cameraTransform.LocalScale.X
	targetY := playerY - (f.screenHeight/2.0)/cameraTransform.LocalScale.Y

	cameraTransform.LocalPosition.X += (targetX - cameraTransform.LocalPosition.X) * smoothingFactor
	cameraTransform.LocalPosition.Y += (targetY - cameraTransform.LocalPosition.Y) * smoothingFactor

	// Ensure the camera stops when close to the target to prevent small jitters
	if abs(cameraTransform.LocalPosition.X-targetX) < 0.2 {
		cameraTransform.LocalPosition.X = targetX
	}
	if abs(cameraTransform.LocalPosition.Y-targetY) < 0.2 {
		cameraTransform.LocalPosition.Y = targetY
	}
}

func (f *CameraFollow) findPlayer(e *ecs.ECS) {
	if f.player == nil {
		entry, found := component.PlayerQuery.First(e.World)
		if found {
			f.player = entry
		}
	}
}

func (f *CameraFollow) findCamera(e *ecs.ECS) {
	if f.camera == nil {
		entry, ok := component.CameraQuery.First(e.World)
		if ok {
			f.camera = entry
		}
	}
}

func (f *CameraFollow) findWorldMap(e *ecs.ECS) {
	if f.worldMap == nil {
		entry, ok := component.WorldMapQuery.First(e.World)
		if ok {
			f.worldMap = entry
		}
	}
}
