package widgets

import (
	"ecs-test/assets/sprites"
	"github.com/tinne26/etxt"
	"github.com/yohamta/furex/v2"
)

type MenuEntry struct {
	Text     string
	Selected bool

	view           *furex.View
	animatedSprite *AnimatedSprite
}

func NewMenuEntry(text string, renderer *etxt.Renderer) *MenuEntry {
	animatedSprite := NewAnimatedSprite(sprites.NewSkullAnimation(), 2)

	compositeView := &furex.View{
		Direction:    furex.Row,
		AlignContent: furex.AlignContentCenter,
		AlignItems:   furex.AlignItemCenter,
	}
	compositeView.AddChild(
		&furex.View{
			Width:   80,
			Height:  60,
			Handler: animatedSprite,
		})
	compositeView.AddChild(
		&furex.View{
			Width: 10,
		})

	compositeView.AddChild(
		&furex.View{
			Width:   250,
			Height:  60,
			Handler: NewText(text, renderer, 2.0).WithVerticalShift(-4).WithShadow(),
		},
	)

	return &MenuEntry{
		Text:           text,
		view:           compositeView,
		animatedSprite: animatedSprite,
	}
}

func (m *MenuEntry) SetSelected(selected bool) *MenuEntry {
	m.Selected = selected
	if selected {
		m.animatedSprite.Start()
	} else {
		m.animatedSprite.Stop()
	}
	return m
}

func (m *MenuEntry) View() *furex.View {
	return m.view
}
