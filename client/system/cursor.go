package system

import (
	"ecs-test/client/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/ganim8/v2"
	"image/color"
	"log"
)

type CursorSystem struct {
	cursorQuery *donburi.Query
	debugQuery  *donburi.Query
	cursorEntry *donburi.Entry
	debugEntry  *donburi.Entry

	mouseDown bool
}

func NewCursorSystem() *CursorSystem {
	return &CursorSystem{
		cursorQuery: donburi.NewQuery(filter.Contains(component.Cursor, transform.Transform)),
		debugQuery:  donburi.NewQuery(filter.Contains(component.Debug)),
	}
}

func (d *CursorSystem) Update(ecs *ecs.ECS) {
	d.findDebugComponent(ecs)
	d.findCursor(ecs)
	d.mouseDown = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	mouseX, mouseY := ebiten.CursorPosition()
	cursorTransform := transform.Transform.Get(d.cursorEntry)
	cursorTransform.LocalPosition.X = float64(mouseX)
	cursorTransform.LocalPosition.Y = float64(mouseY)
}

func (d *CursorSystem) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	if d.cursorEntry == nil {
		return
	}

	cursorData := component.Cursor.Get(d.cursorEntry)
	cursorTransform := transform.Transform.Get(d.cursorEntry)

	drawOpts := ganim8.DrawOpts(cursorTransform.LocalPosition.X, cursorTransform.LocalPosition.Y, 0, 2, 2)
	if d.mouseDown {
		d.drawPressedCursor(cursorData, screen, drawOpts)
	} else {
		d.drawNormalCursor(cursorData, screen, drawOpts)
	}

	debug := component.Debug.Get(d.debugEntry)
	if debug.IsEnabled() {
		vector.DrawFilledCircle(screen, float32(cursorTransform.LocalPosition.X), float32(cursorTransform.LocalPosition.Y), 5, color.RGBA{G: 255}, false)
	}
}

func (d *CursorSystem) drawNormalCursor(cursorData *component.CursorData, screen *ebiten.Image, drawOpts *ganim8.DrawOptions) {
	cursorData.Sprite.Draw(screen, 0, drawOpts)
	drawOpts.SetPos(drawOpts.X+float64(cursorData.Sprite.W())*drawOpts.ScaleX, drawOpts.Y)
	cursorData.Sprite.Draw(screen, 1, drawOpts)
}

func (d *CursorSystem) drawPressedCursor(cursorData *component.CursorData, screen *ebiten.Image, drawOpts *ganim8.DrawOptions) {
	cursorData.Sprite.Draw(screen, 4, drawOpts)
	drawOpts.SetPos(drawOpts.X+float64(cursorData.Sprite.W())*drawOpts.ScaleX, drawOpts.Y)
	cursorData.Sprite.Draw(screen, 5, drawOpts)
}

func (d *CursorSystem) findCursor(e *ecs.ECS) {
	if d.cursorEntry == nil {
		cursorEntry, entryFound := d.cursorQuery.First(e.World)
		if !entryFound {
			log.Fatalf("Cursor entry not found!")
			return
		}
		d.cursorEntry = cursorEntry
	}
}

func (d *CursorSystem) findDebugComponent(e *ecs.ECS) {
	if d.debugEntry == nil {
		debugEntry, entryFound := d.debugQuery.First(e.World)
		if !entryFound {
			log.Fatalf("Debug entry not found!")
			return
		}
		d.debugEntry = debugEntry
	}
}
