package component

import (
	"ecs-test/assets"
	"ecs-test/client/camera"
	"ecs-test/client/component/world"
	"ecs-test/client/loader"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"log"
	"time"
)

type WorldMapData struct {
	Loaded        bool
	layers        map[world.LayerID]*world.Layer
	WalkableLayer *world.WalkableLayer

	Width, Height         int
	tileWidth, tileHeight int
}

func (w *WorldMapData) LoadFromTiledFile(fileLocation string, monitor *loader.ProgressMonitor) {
	worldMap, err := tiled.LoadFile(fileLocation, tiled.WithFileSystem(assets.ScenesAssets))
	if err != nil {
		log.Panic(err)
	}
	w.Width = worldMap.Width * worldMap.TileWidth    // Convert to pixels
	w.Height = worldMap.Height * worldMap.TileHeight // Convert to pixels

	renderer := world.NewLayerRenderer()

	w.layers = map[world.LayerID]*world.Layer{
		world.Ground: loader.Progress(monitor, func() *world.Layer {
			return renderer.LoadWorldLayer(worldMap, world.Ground)
		}),
		world.Elevation: loader.Progress(monitor, func() *world.Layer {
			return renderer.LoadWorldLayer(worldMap, world.Elevation)
		}),
		world.Decorations: loader.Progress(monitor, func() *world.Layer {
			return renderer.LoadWorldLayer(worldMap, world.Decorations)
		}),
		world.OtherDecorations: loader.Progress(monitor, func() *world.Layer {
			return renderer.LoadWorldLayer(worldMap, world.OtherDecorations)
		}),
	}

	walkableReader := world.NewWalkableLayerReader()
	w.WalkableLayer = loader.Progress(monitor, func() *world.WalkableLayer {
		return walkableReader.LoadWalkableLayer(worldMap, world.Walkable)
	})

	w.Loaded = true
	renderer.Cleanup()
	renderer = nil
}

func (w *WorldMapData) IsLoaded() bool {
	return w.Loaded
}

func (w *WorldMapData) DrawLayer(
	layerId world.LayerID,
	screen *ebiten.Image,
	options *ebiten.DrawImageOptions,
	viewport camera.Viewport,
) {
	w.forEachTileInViewport(layerId, viewport, func(tile world.Tile) {
		tile.Draw(screen, options)
	})
}

func (w *WorldMapData) DrawWalkableLayer(
	screen *ebiten.Image,
	options *ebiten.DrawImageOptions,
	viewport camera.Viewport,
) {
	w.WalkableLayer.Draw(screen, options, viewport)
}

func (w *WorldMapData) UpdateLayer(layerId world.LayerID, viewport camera.Viewport, deltaTime time.Duration) {
	w.forEachTileInParallel(layerId, func(tile world.Tile) {
		tile.Update(deltaTime)
	})
}

func (w *WorldMapData) forEachTileInViewport(layerId world.LayerID, viewport camera.Viewport, action func(tile world.Tile)) {
	layer, ok := w.layers[layerId]
	if !ok || layer == nil {
		return
	}

	layer.ForEachTileInViewportOptimized(viewport, action)
	//layer.ForEachTileInViewport(viewport, action)
	//layer.ForEachTile(action)
}

func (w *WorldMapData) forEachTileInParallel(layerId world.LayerID, action func(tile world.Tile)) {
	layer, ok := w.layers[layerId]
	if !ok || layer == nil {
		return
	}

	layer.ForEachTileParallel(action, 8)
}

var WorldMap = donburi.NewComponentType[WorldMapData]()
var WorldMapQuery = donburi.NewQuery(filter.Contains(WorldMap))
