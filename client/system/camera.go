package system

import (
	"ecs-test/client/component"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type Camera struct {
	camera          *donburi.Entry
	targetZoomScale dmath.Vec2
	minScale        dmath.Vec2
	screenWidth     float64
	screenHeight    float64
	speed           float64
	zoomTransition  bool
}

func NewCamera(
	screenWidth, screenHeight int,
	initialScale float64,
	speed float64,
) *Camera {
	return &Camera{
		targetZoomScale: dmath.Vec2{X: initialScale, Y: initialScale}, // Initial target scale
		minScale:        dmath.Vec2{X: initialScale, Y: initialScale}, // Minimum scale limit
		zoomTransition:  false,
		screenWidth:     float64(screenWidth),
		screenHeight:    float64(screenHeight),
		speed:           speed,
	}
}

func (c *Camera) Update(e *ecs.ECS) {
	c.findCamera(e)

	if c.camera == nil {
		return
	}

	transformData := transform.Transform.Get(c.camera)
	cameraData := component.Camera.Get(c.camera)

	c.updateCameraPosition(transformData, cameraData)
	c.updateCameraZoom(cameraData)

	if c.zoomTransition {
		oldScaleX := transformData.LocalScale.X
		oldScaleY := transformData.LocalScale.Y
		oldPosX := transformData.LocalPosition.X
		oldPosY := transformData.LocalPosition.Y

		screenCenterX := (c.screenWidth / 2 / oldScaleX) + oldPosX
		screenCenterY := (c.screenHeight / 2 / oldScaleY) + oldPosY

		transformData.LocalScale.X += (c.targetZoomScale.X - oldScaleX) * cameraData.ZoomSpeed
		transformData.LocalScale.Y += (c.targetZoomScale.Y - oldScaleY) * cameraData.ZoomSpeed

		newPosX := screenCenterX - (c.screenWidth / 2 / transformData.LocalScale.X)
		newPosY := screenCenterY - (c.screenHeight / 2 / transformData.LocalScale.Y)

		transformData.LocalPosition.X = newPosX
		transformData.LocalPosition.Y = newPosY

		// Stop the zoom transition when the target scale is reached
		if abs(c.targetZoomScale.X-transformData.LocalScale.X) < 0.001 &&
			abs(c.targetZoomScale.Y-transformData.LocalScale.Y) < 0.001 {
			transformData.LocalScale = c.targetZoomScale
			c.zoomTransition = false
		}
	} else if transformData.LocalScale != c.targetZoomScale {
		transformData.LocalScale = c.targetZoomScale
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		fmt.Printf("%f %f\n", transformData.LocalPosition.X, transformData.LocalPosition.Y)
	}

	cameraData.UpdateViewport(transformData, dmath.NewVec2(c.screenWidth, c.screenHeight))
}

func (c *Camera) updateCameraZoom(cameraData *component.CameraData) {
	if inpututil.IsKeyJustPressed(ebiten.KeyPageUp) {
		c.targetZoomScale.X *= 1 + cameraData.ZoomFactor
		c.targetZoomScale.Y *= 1 + cameraData.ZoomFactor
		c.zoomTransition = true
	} else if inpututil.IsKeyJustPressed(ebiten.KeyPageDown) {
		newScaleX := c.targetZoomScale.X * (1.0 - cameraData.ZoomFactor)
		newScaleY := c.targetZoomScale.Y * (1.0 - cameraData.ZoomFactor)

		if newScaleX >= c.minScale.X && newScaleY >= c.minScale.Y {
			c.targetZoomScale.X = newScaleX
			c.targetZoomScale.Y = newScaleY
		} else {
			c.targetZoomScale = c.minScale
		}

		c.zoomTransition = true
	}
}

func (c *Camera) updateCameraPosition(
	transformData *transform.TransformData,
	cameraData *component.CameraData,
) {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		transformData.LocalPosition.Y -= c.speed
		cameraData.FreeLook()
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		transformData.LocalPosition.Y += c.speed
		cameraData.FreeLook()
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		transformData.LocalPosition.X -= c.speed
		cameraData.FreeLook()
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		transformData.LocalPosition.X += c.speed
		cameraData.FreeLook()
	}
}

func (c *Camera) findCamera(e *ecs.ECS) {
	if c.camera == nil {
		entry, ok := component.CameraQuery.First(e.World)
		if ok {
			c.camera = entry
		}
	}
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}
