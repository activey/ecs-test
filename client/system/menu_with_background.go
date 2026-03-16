package system

import (
	"ecs-test/client/view/effects"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type MenuWithBackground struct {
	menu       *Menu
	background *GifRender
	fadeToGray *effects.FadeToGray
}

func NewMenuWithBackground(world donburi.World, width, height int) *MenuWithBackground {
	menu := NewMenu(width, height, world)
	background, err := NewGifRender("assets/background.gif", 1.0)
	if err != nil {
		panic(err)
	}

	return &MenuWithBackground{
		menu:       menu,
		background: background,
		fadeToGray: effects.NewFadeToGray(width, height),
	}
}

func (m *MenuWithBackground) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	m.background.Draw(screen)
	m.fadeToGray.Draw(screen, screen)
	m.menu.Draw(screen)
}

func (m *MenuWithBackground) Update(ecs *ecs.ECS) {
	m.background.Update(ecs)
	m.menu.Update(ecs)
	m.fadeToGray.Update()
}

func (m *MenuWithBackground) FadeToGray() {
	m.fadeToGray.FadeToGray()
}

func (m *MenuWithBackground) FadeToNormal() {
	m.fadeToGray.FadeToNormal()
}
