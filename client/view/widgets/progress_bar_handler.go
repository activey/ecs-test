package widgets

import (
	"ecs-test/assets/sprites/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/ganim8/v2"
	"image"
)

type ProgressBarHandler struct {
}

func NewProgressBarHandler() *ProgressBarHandler {
	return &ProgressBarHandler{}
}

func (p ProgressBarHandler) Draw(screen *ebiten.Image, frame image.Rectangle, v *furex.View) {
	sprite := ui.ProgressBarSprite
	horizontalFrames := 11
	frameSize := 16

	offscreen := ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy())

	horizontal, vertical := 0, 0
	for i := 0; i < sprite.Length(); i++ {
		sprite.Draw(
			offscreen,
			i,
			ganim8.DrawOpts(
				float64(horizontal*frameSize),
				float64(vertical*frameSize),
			),
		)

		horizontal++
		if horizontal >= horizontalFrames {
			horizontal = 0
			vertical++
		}
	}

	// Scaling factor
	scaleFactor := 2.0
	totalWidth := float64(horizontalFrames * frameSize * int(scaleFactor)) // Calculate total width after scaling
	totalHeight := float64((vertical + 1) * frameSize * int(scaleFactor))  // Calculate total height after scaling

	// Calculate the horizontal offset to center the image
	centerX := float64(frame.Min.X) + (float64(frame.Dx())-totalWidth)/2
	centerY := float64(frame.Min.Y) + (float64(frame.Dy())-totalHeight)/2

	// Set up drawing options for scaling and translation
	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Scale(scaleFactor, scaleFactor)
	ops.GeoM.Translate(centerX, centerY)

	// Draw the offscreen image onto the screen
	screen.DrawImage(offscreen, ops)
}
