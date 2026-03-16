package effects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

type FadeToGray struct {
	timer        float64
	fadingToGray bool

	offscreen *ebiten.Image
}

func NewFadeToGray(screenWidth, screenHeight int) *FadeToGray {
	return &FadeToGray{
		offscreen: ebiten.NewImage(screenWidth, screenHeight),
	}
}

func (g *FadeToGray) Update() {
	changeSpeed := 2.0
	dt := 1.0 / float64(ebiten.TPS()) * changeSpeed
	if g.fadingToGray {
		g.timer += dt
		if g.timer >= 1.0 {
			g.timer = 1.0
		}
	} else {
		g.timer -= dt
		if g.timer < 0 {
			g.timer = 0
		}
	}
}

func (g *FadeToGray) Draw(source *ebiten.Image, screen *ebiten.Image) {
	//g.offscreen.Clear()
	g.offscreen.DrawImage(source, nil)

	colorM := colorm.ColorM{}
	saturation := g.lerp(1.0, 0.0, g.timer)
	colorM.ChangeHSV(0, saturation, 1)
	colorm.DrawImage(screen, g.offscreen, colorM, &colorm.DrawImageOptions{})
}

func (g *FadeToGray) FadeToGray() {
	g.fadingToGray = true
}

func (g *FadeToGray) FadeToNormal() {
	g.fadingToGray = false
}

func (g *FadeToGray) lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}
