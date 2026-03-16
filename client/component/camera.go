package component

import (
	"ecs-test/client/camera"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type CameraData struct {
	ZoomSpeed       float64
	ZoomFactor      float64
	VisibleViewport camera.Viewport

	MovementSpeed float64
	followPlayer  bool
}

func (d *CameraData) FollowPlayer() {
	d.followPlayer = true
}

func (d *CameraData) FreeLook() {
	d.followPlayer = false
}

func (d *CameraData) ShouldFollowPlayer() bool {
	return d.followPlayer
}

func (d *CameraData) UpdateViewport(
	cameraTransform *transform.TransformData,
	screenDimensions dmath.Vec2,
) {
	minX := cameraTransform.LocalPosition.X
	minY := cameraTransform.LocalPosition.Y
	maxX := minX + (screenDimensions.X / cameraTransform.LocalScale.X)
	maxY := minY + (screenDimensions.Y / cameraTransform.LocalScale.Y)

	d.VisibleViewport = camera.Viewport{
		Min: dmath.NewVec2(minX, minY),
		Max: dmath.NewVec2(maxX, maxY),
	}
}

var Camera = donburi.NewComponentType[CameraData]()
var CameraQuery = donburi.NewQuery(filter.Contains(Camera, transform.Transform))
