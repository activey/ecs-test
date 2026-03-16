package widgets

import (
	"ecs-test/assets/sprites/ui"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/ganim8/v2"
	"image"
)

type ProgressBarFillHandler struct {
	progress int
}

func NewProgressBarFillHandler(int) *ProgressBarFillHandler {
	return &ProgressBarFillHandler{}
}

func (p *ProgressBarFillHandler) Draw(screen *ebiten.Image, frame image.Rectangle, v *furex.View) {
	sprite := ui.ProgressBarFillerSprite
	horizontalFrames := 11
	frameSize := 16

	// Create an offscreen image with the scaled size
	scaleFactor := 2.0
	offscreenWidth := screen.Bounds().Dx()
	offscreenHeight := screen.Bounds().Dy()
	offscreen := ebiten.NewImage(offscreenWidth, offscreenHeight)

	// Draw the entire sprite scaled up onto the offscreen buffer
	horizontal, vertical := 0, 0
	for i := 0; i < sprite.Length(); i++ {
		ops := ganim8.DrawOpts(float64(horizontal*frameSize)*scaleFactor, float64(vertical*frameSize)*scaleFactor, 0, scaleFactor, scaleFactor)

		sprite.Draw(offscreen, i, ops)

		horizontal++
		if horizontal >= horizontalFrames {
			horizontal = 0
			vertical++
		}
	}

	// Now take a portion based on the progress and scaled image
	progressFactor := float64(p.progress) / 100.0
	scaledWidth := frame.Dx()
	croppedWidth := int(float64(scaledWidth) * progressFactor)

	if croppedWidth == 0 {
		return
	}

	// SubImage from the scaled offscreen image
	subRect := image.Rect(0, 0, croppedWidth, offscreenHeight)
	subImage := ebiten.NewImageFromImage(offscreen.SubImage(subRect))

	// Draw the cropped portion back to the screen
	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Translate(float64(frame.Min.X), float64(frame.Min.Y))

	// Draw the sub-image onto the screen
	screen.DrawImage(subImage, ops)
}

func (p *ProgressBarFillHandler) Update(v *furex.View) {
}

func (p *ProgressBarFillHandler) UpdateProgress(progress int) {
	p.progress = progress
	progressFactor := float64(p.progress) / 100.0

	fmt.Printf("progress: %f\n", progressFactor)
}
