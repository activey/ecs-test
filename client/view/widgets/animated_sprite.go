package widgets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/ganim8/v2"
	"image"
)

type AnimatedSprite struct {
	animation *ganim8.Animation

	animating bool
	scale     float64
}

func NewAnimatedSprite(animation *ganim8.Animation, scale float64) *AnimatedSprite {
	return &AnimatedSprite{
		animation: animation,
		scale:     scale,
	}
}

func (a *AnimatedSprite) Update(v *furex.View) {
	if !a.animating {
		return
	}
	a.animation.Update()
}

func (a *AnimatedSprite) Draw(screen *ebiten.Image, frame image.Rectangle, v *furex.View) {
	if !a.animating {
		return
	}

	frameWidth, frameHeight := a.animation.Size()
	containerHeight := float64(frame.Dy())
	containerWidth := float64(frame.Dx())

	spriteHeight := float64(frameHeight) * a.scale
	spriteWidth := float64(frameWidth) * a.scale

	x := float64(frame.Min.X) + (containerWidth - spriteWidth)
	y := float64(frame.Min.Y) + (containerHeight-spriteHeight)/2

	a.animation.Draw(
		screen,
		ganim8.DrawOpts(
			x,
			y,
			0,
			a.scale,
			a.scale,
		),
	)
}

func (a *AnimatedSprite) Stop() {
	a.animating = false
}

func (a *AnimatedSprite) Start() {
	a.animating = true
}
