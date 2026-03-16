package system

import (
	"ecs-test/client/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type CameraBoundary struct {
	worldMapQuery *query.Query
	screenWidth   float64
	screenHeight  float64

	worldMapEntity *donburi.Entry
	cameraEntity   *donburi.Entry
}

func NewCameraBoundary(screenWidth, screenHeight int) *CameraBoundary {
	return &CameraBoundary{
		worldMapQuery: query.NewQuery(filter.Contains(component.WorldMap)),
		screenWidth:   float64(screenWidth),
		screenHeight:  float64(screenHeight),
	}
}

func (cb *CameraBoundary) Update(e *ecs.ECS) {
	cb.findWorldMap(e)
	cb.findCamera(e)

	worldMapData := component.WorldMap.Get(cb.worldMapEntity)
	cameraTransform := transform.Transform.Get(cb.cameraEntity)

	// Get the current camera scale (zoom level)
	cameraScaleX := cameraTransform.LocalScale.X
	cameraScaleY := cameraTransform.LocalScale.Y

	// Define the camera viewport size based on the screen size and scale
	cameraWidth := cb.screenWidth / cameraScaleX
	cameraHeight := cb.screenHeight / cameraScaleY

	// Define the minimum and maximum positions based on world size and scaled camera size
	minX := 0.0
	minY := 0.0
	maxX := float64(worldMapData.Width) - cameraWidth
	maxY := float64(worldMapData.Height) - cameraHeight

	// Clamp the camera position to stay within bounds
	cameraTransform.LocalPosition.X = clamp(cameraTransform.LocalPosition.X, minX, maxX)
	cameraTransform.LocalPosition.Y = clamp(cameraTransform.LocalPosition.Y, minY, maxY)

	cameraData := component.Camera.Get(cb.cameraEntity)
	cameraData.UpdateViewport(cameraTransform, math.NewVec2(cb.screenWidth, cb.screenHeight))
}

func (cb *CameraBoundary) findWorldMap(e *ecs.ECS) {
	if cb.worldMapEntity == nil {
		worldMapEntity, worldMapFound := cb.worldMapQuery.First(e.World)
		if !worldMapFound {
			// Handle case where the world map entity is not found
		}
		cb.worldMapEntity = worldMapEntity
	}
}

func (cb *CameraBoundary) findCamera(e *ecs.ECS) {
	if cb.cameraEntity == nil {
		cameraEntity, cameraFound := component.CameraQuery.First(e.World)
		if !cameraFound {
			// Handle case where the world map entity is not found
		}
		cb.cameraEntity = cameraEntity
	}
}

// Helper function to clamp values
func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}
