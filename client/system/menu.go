package system

import (
	"ecs-test/assets/fonts"
	"ecs-test/client/event"
	"ecs-test/client/view/widgets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tinne26/etxt"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/furex/v2"
	"image/color"
)

type Menu struct {
	ui       *furex.View
	menuView *furex.View

	selectedIdx int
	menuEntries []*widgets.MenuEntry
	world       donburi.World
	offscreen   *ebiten.Image
}

func NewMenu(width, height int, w donburi.World) *Menu {
	rootView := &furex.View{
		ID:           "root",
		Width:        width,
		Height:       height,
		Direction:    furex.Row,
		AlignContent: furex.AlignContentCenter,
		Justify:      furex.JustifyCenter,
		AlignItems:   furex.AlignItemCenter,
	}

	menuView := &furex.View{
		ID:           "menu",
		Height:       250,
		Direction:    furex.Column,
		AlignContent: furex.AlignContentCenter,
		Justify:      furex.JustifyCenter,
		AlignItems:   furex.AlignItemCenter,
	}

	r := etxt.NewRenderer()
	r.SetFont(fonts.MainFont)
	r.SetSize(72)
	r.SetColor(color.RGBA{
		R: 217,
		G: 179,
		B: 132,
		A: 255,
	})

	menuEntries := []*widgets.MenuEntry{
		widgets.NewMenuEntry("Join World", r).SetSelected(true),
		widgets.NewMenuEntry("Settings", r),
		widgets.NewMenuEntry("Quit", r),
	}

	for _, entry := range menuEntries {
		menuView.AddChild(entry.View())
	}
	menuEntries[0].SetSelected(true)

	rootView.AddChild(menuView)
	return &Menu{
		world:       w,
		ui:          rootView,
		menuView:    menuView,
		menuEntries: menuEntries,

		offscreen: ebiten.NewImage(width, height),
	}
}

func (m *Menu) Draw(screen *ebiten.Image) {
	m.ui.Draw(screen)
}

func (m *Menu) Update(ecs *ecs.ECS) {
	m.updateMenuSelection()
	m.ui.Update()
}

func (m *Menu) moveSelection(direction int) {
	m.menuEntries[m.selectedIdx].SetSelected(false)
	m.selectedIdx = (m.selectedIdx + direction + len(m.menuEntries)) % len(m.menuEntries)
	m.menuEntries[m.selectedIdx].SetSelected(true)
}

func (m *Menu) updateMenuSelection() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		m.moveSelection(1)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		m.moveSelection(-1)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		m.executeAction()
	}
}

func (m *Menu) executeAction() {
	switch m.selectedIdx {
	case 0:
		m.joinWorld()
	case 2:
		m.quit()
	}
}

func (m *Menu) joinWorld() {
	event.MenuSelectionEvent.Publish(m.world, event.JoinWorld)
}

func (m *Menu) quit() {
	println("quitting")
	event.MenuSelectionEvent.Publish(m.world, event.Quit)
}
