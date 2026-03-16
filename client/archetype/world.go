package archetype

import (
	"ecs-test/client/component"
	"github.com/yohamta/donburi"
)

func NewWorldMap(w donburi.World) *donburi.Entry {
	worldMap := w.Entry(
		w.Create(
			component.WorldMap,
		),
	)

	// The world map can be loaded later using the `LoadFromTiledFile` method
	worldMapData := component.WorldMap.Get(worldMap)
	worldMapData.Loaded = false // Set it as not loaded initially

	return worldMap
}
