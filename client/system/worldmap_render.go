package system

import (
	"ecs-test/client/component"
	"ecs-test/client/component/world"
	"ecs-test/client/view/effects"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"log"
)

type WorldMapRender struct {
	worldMap   *donburi.Entry
	camera     *donburi.Entry
	debug      *donburi.Entry
	debugQuery *donburi.Query
	shutter    *effects.Shutter

	worldRevealed  bool
	shutterRunning bool
	offScreens     map[world.LayerID]*ebiten.Image
}

func NewWorldMapRender(
	screenWidth int,
	screenHeight int,
) *WorldMapRender {

	w := &WorldMapRender{
		debugQuery:     donburi.NewQuery(filter.Contains(component.Debug)),
		shutterRunning: true,

		offScreens: map[world.LayerID]*ebiten.Image{
			world.Ground:           ebiten.NewImage(screenWidth, screenHeight),
			world.Elevation:        ebiten.NewImage(screenWidth, screenHeight),
			world.Decorations:      ebiten.NewImage(screenWidth, screenHeight),
			world.OtherDecorations: ebiten.NewImage(screenWidth, screenHeight),
		},
	}

	shutter := effects.NewShutter(screenWidth, screenHeight, 200, 1)
	shutter.Start(effects.ShutterExpand, func() {
		w.shutterRunning = false
	})

	w.shutter = shutter
	return w
}

func (w *WorldMapRender) drawLayer(layerId world.LayerID, screen *ebiten.Image) {
	if w.worldMap == nil {
		return
	}

	worldMap := component.WorldMap.Get(w.worldMap)
	if !worldMap.IsLoaded() {
		return
	}

	cameraViewport := component.Camera.Get(w.camera).VisibleViewport

	offScreen := w.offScreens[layerId]
	offScreen.Clear()
	worldMap.DrawLayer(layerId, offScreen, w.drawOptions(), cameraViewport)

	screen.DrawImage(offScreen, nil)
}

func (w *WorldMapRender) DrawGround(ecs *ecs.ECS, screen *ebiten.Image) {
	w.drawLayer(world.Ground, screen)
}

func (w *WorldMapRender) DrawElevation(ecs *ecs.ECS, screen *ebiten.Image) {
	w.drawLayer(world.Elevation, screen)
}

func (w *WorldMapRender) DrawDecorations(ecs *ecs.ECS, screen *ebiten.Image) {
	w.drawLayer(world.Decorations, screen)
}

func (w *WorldMapRender) DrawOtherDecorations(ecs *ecs.ECS, screen *ebiten.Image) {
	w.drawLayer(world.OtherDecorations, screen)
}

func (w *WorldMapRender) DrawWalkable(ecs *ecs.ECS, screen *ebiten.Image) {
	if w.debug == nil || w.worldMap == nil {
		return
	}
	debugData := component.Debug.Get(w.debug)
	worldMap := component.WorldMap.Get(w.worldMap)
	cameraData := component.Camera.Get(w.camera)

	if debugData.IsEnabled() && worldMap.IsLoaded() {
		worldMap.DrawWalkableLayer(screen, w.drawOptions(), cameraData.VisibleViewport)
	}
}

func (w *WorldMapRender) drawOptions() *ebiten.DrawImageOptions {
	cameraTransform := transform.Transform.Get(w.camera)
	positionX, positionY := cameraTransform.LocalPosition.XY()

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(-positionX, -positionY)
	options.GeoM.Scale(cameraTransform.LocalScale.X, cameraTransform.LocalScale.Y)

	return options
}

func (w *WorldMapRender) DrawShutter(ecs *ecs.ECS, screen *ebiten.Image) {
	if w.worldMap == nil {
		return
	}
	worldMap := component.WorldMap.Get(w.worldMap)
	if !worldMap.IsLoaded() {
		return
	}
	w.shutter.Draw(screen)
}

func (w *WorldMapRender) Update(e *ecs.ECS) {
	w.findDebugComponent(e)
	w.findCamera(e)
	w.findWorldMap(e)

	if w.worldMap == nil {
		return
	}

	worldMap := component.WorldMap.Get(w.worldMap)
	if !worldMap.IsLoaded() {
		return
	}

	w.updateShutter(e)

	deltaTime := e.Time.DeltaTime()
	cameraViewport := component.Camera.Get(w.camera).VisibleViewport
	worldMap.UpdateLayer(world.Ground, cameraViewport, deltaTime)
	worldMap.UpdateLayer(world.Elevation, cameraViewport, deltaTime)
	worldMap.UpdateLayer(world.Decorations, cameraViewport, deltaTime)
	worldMap.UpdateLayer(world.OtherDecorations, cameraViewport, deltaTime)
}

func (w *WorldMapRender) updateShutter(e *ecs.ECS) {
	if w.shutterRunning {
		w.shutter.Update()
	}
}

func (w *WorldMapRender) findDebugComponent(e *ecs.ECS) {
	if w.debug == nil {
		debugEntry, entryFound := w.debugQuery.First(e.World)
		if !entryFound {
			log.Fatalf("Debug entry not found!")
			return
		}
		w.debug = debugEntry
	}
}

func (w *WorldMapRender) findCamera(e *ecs.ECS) {
	if w.camera == nil {
		cameraEntry, found := component.CameraQuery.First(e.World)
		if !found {
		}
		w.camera = cameraEntry
	}
}

func (w *WorldMapRender) findWorldMap(e *ecs.ECS) {
	if w.worldMap == nil {
		entry, ok := component.WorldMapQuery.First(e.World)
		if ok {
			w.worldMap = entry
		}
	}
}
