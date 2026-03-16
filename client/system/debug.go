package system

import (
	"ecs-test/client/component"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
	"log"
)

type DebugSystem struct {
	debugQuery *donburi.Query
	debugEntry *donburi.Entry
}

func NewDebugSystem() *DebugSystem {
	return &DebugSystem{
		debugQuery: donburi.NewQuery(filter.Contains(component.Debug)),
	}
}

func (d *DebugSystem) Update(ecs *ecs.ECS) {
	d.findDebugComponent(ecs)

	debugData := component.Debug.Get(d.debugEntry)
	if inpututil.IsKeyJustPressed(ebiten.KeySlash) {
		debugData.ToggleDebug()
	}
}

func (d *DebugSystem) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	if d.debugEntry == nil {
		// for some reason happens at first frame
		return
	}

	debugData := component.Debug.Get(d.debugEntry)
	if debugData.IsEnabled() {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f TPS: %.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
	}
}

func (d *DebugSystem) findDebugComponent(e *ecs.ECS) {
	if d.debugEntry == nil {
		debugEntry, entryFound := d.debugQuery.First(e.World)
		if !entryFound {
			log.Fatalf("Debug entry not found!")
			return
		}
		d.debugEntry = debugEntry
	}
}
