package widgets

import (
	"ecs-test/assets/sprites/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/ganim8/v2"
	"image"
)

type Header struct {
	scale float64
}

func NewHeader(scale float64) *Header {
	return &Header{
		scale: scale,
	}
}

func (t *Header) Update(view *furex.View) {
	sprite := ui.HeaderSprite
	spriteHeight := float64(sprite.Height())
	scaledFrameHeight := spriteHeight * t.scale
	view.SetHeight(int(scaledFrameHeight))
}

func (t *Header) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	sprite := ui.HeaderSprite

	// Get frame dimensions
	spriteWidth := float64(sprite.Width()) // Width of each sprite frame without scaling

	// Calculate scaled dimensions
	scaledFrameWidth := spriteWidth * t.scale

	// Position to draw the left frame
	leftX := float64(frame.Min.X)
	leftY := float64(frame.Min.Y)

	// Calculate remaining width for the middle frame
	middleWidth := float64(frame.Dx()) - 2*scaledFrameWidth
	if middleWidth < 0 {
		middleWidth = 0 // Ensure non-negative
	}

	// 1. Draw the left frame (index 0)
	leftOpts := ganim8.DrawOpts(leftX, leftY)
	leftOpts.SetScale(t.scale, t.scale)
	sprite.Draw(screen, 0, leftOpts)

	// 2. Draw the middle frame (index 1)
	if middleWidth > 0 {
		middleX := leftX + scaledFrameWidth
		tileCount := middleWidth / scaledFrameWidth

		// Adjust tile width if necessary
		for i := 0.0; i < tileCount; i++ {
			middleOpts := ganim8.DrawOpts(middleX+i*scaledFrameWidth, leftY)
			middleOpts.SetScale(t.scale, t.scale)
			sprite.Draw(screen, 1, middleOpts)
		}

		// Draw any remaining width with a partial tile
		if remainingWidth := middleWidth - (tileCount * scaledFrameWidth); remainingWidth > 0 {
			middleOpts := ganim8.DrawOpts(middleX+tileCount*scaledFrameWidth, leftY, 0, remainingWidth/spriteWidth, t.scale)
			sprite.Draw(screen, 1, middleOpts)
		}
	}

	// 3. Draw the right frame (index 2)
	rightX := leftX + scaledFrameWidth + middleWidth
	rightOpts := ganim8.DrawOpts(rightX, leftY)
	rightOpts.SetScale(t.scale, t.scale)
	sprite.Draw(screen, 2, rightOpts)
}
