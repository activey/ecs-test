package system

import (
	"ecs-test/assets/fonts"
	"ecs-test/client/view/widgets"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/etxt"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/furex/v2"
	"time"
)

type GameUi struct {
	rootView  *furex.View
	clockText *widgets.Text

	elapsedTime   time.Duration
	secondsPassed int64
}

func NewGameUi(screenWidth, screenHeight int) *GameUi {
	r := etxt.NewRenderer()
	r.SetFont(fonts.MainFont)
	r.SetSize(52)

	rootView := &furex.View{
		Width:        screenWidth,
		Height:       screenHeight,
		Direction:    furex.Row,
		AlignContent: furex.AlignContentEnd,
		Justify:      furex.JustifyEnd,
		AlignItems:   furex.AlignItemStart,
	}

	clockText := widgets.NewText("Test", r, 2.0).WithAlignment(widgets.RightAlign).WithShadow()
	rootView.AddChild(
		&furex.View{
			Width:       100,
			Height:      50,
			Handler:     clockText,
			MarginRight: 5,
		},
	)

	return &GameUi{
		rootView:  rootView,
		clockText: clockText,
	}
}

func (g *GameUi) Update(ecs *ecs.ECS) {
	g.elapsedTime += ecs.Time.DeltaTime()

	if g.elapsedTime >= time.Second {
		g.secondsPassed++
		g.elapsedTime -= time.Second // Reset elapsed time while keeping any extra fraction of a second
	}

	g.rootView.Update()
	g.clockText.SetText(fmt.Sprintf("%d", g.secondsPassed))
}

func (g *GameUi) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	g.rootView.Draw(screen)
}
