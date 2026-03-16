package archetype

import (
	"ecs-test/client/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func NewCamera(
	w donburi.World,
	startPosition math.Vec2,
	zoomSpeed float64,
	zoomFactor float64,
) *donburi.Entry {
	camera := w.Entry(
		w.Create(
			transform.Transform,
			component.Camera,
		),
	)

	// camera setup
	cameraData := component.Camera.Get(camera)
	cameraData.ZoomSpeed = zoomSpeed
	cameraData.ZoomFactor = zoomFactor
	cameraData.FollowPlayer()

	// position setup
	transform.Transform.Get(camera).LocalPosition = startPosition

	return camera
}
