package system

import (
	"ecs-test/assets/fonts"
	"ecs-test/client/component"
	"ecs-test/client/loader"
	"ecs-test/client/view/widgets"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/etxt"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/furex/v2"
)

type WorldMapLoader struct {
	ui               *furex.View
	worldMap         *donburi.Entry
	progressBar      *widgets.ProgressBar
	currentlyLoading bool
}

func NewWorldMapLoader(
	width, height int,
) *WorldMapLoader {
	rootView := &furex.View{
		Width:        width,
		Height:       height,
		Direction:    furex.Column,
		AlignContent: furex.AlignContentCenter,
		Justify:      furex.JustifyEnd,
		AlignItems:   furex.AlignItemCenter,
	}

	progressBar := widgets.NewProgressBar(0)
	rootView.AddChild(
		progressBar.View(),
	)

	r := etxt.NewRenderer()
	r.SetFont(fonts.MainFont)
	r.SetSize(22)

	rootView.AddChild(
		&furex.View{
			Width:   100,
			Height:  30,
			Handler: widgets.NewText("Generating world ...", r, 1.0).WithAlignment(widgets.CenterAlign),
		},
	)

	return &WorldMapLoader{
		ui:          rootView,
		progressBar: progressBar,
	}
}

func (w *WorldMapLoader) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	if w.worldMap == nil {
		return
	}
	worldMap := component.WorldMap.Get(w.worldMap)
	if worldMap.IsLoaded() {
		return
	}

	w.ui.Draw(screen)
}

func (w *WorldMapLoader) Update(e *ecs.ECS) {
	w.findWorldMap(e)
	if w.worldMap == nil {
		return
	}

	worldMap := component.WorldMap.Get(w.worldMap)
	if worldMap.IsLoaded() || w.currentlyLoading {
		return
	}

	w.currentlyLoading = true
	go w.load()
}

func (w *WorldMapLoader) load() {
	fmt.Println("loading world")
	progressMonitor := loader.NewProgressMonitor(5, w.progressBar.UpdateProgress)

	worldMap := component.WorldMap.Get(w.worldMap)
	worldMap.LoadFromTiledFile("scenes/world2.tmx", progressMonitor)
	w.currentlyLoading = false
}

func (w *WorldMapLoader) findWorldMap(e *ecs.ECS) {
	if w.worldMap == nil {
		entry, ok := component.WorldMapQuery.First(e.World)
		if ok {
			w.worldMap = entry
		}
	}
}
