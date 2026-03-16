package widgets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
)

type TextRenderer struct {
	fontSprite              *ganim8.Sprite
	frameWidth, frameHeight float64
}

func NewTextRenderer(spriteSheet *ganim8.Sprite, frameWidth, frameHeight float64) *TextRenderer {
	return &TextRenderer{
		fontSprite:  spriteSheet,
		frameWidth:  frameWidth,
		frameHeight: frameHeight,
	}
}

func (t *TextRenderer) RenderTextWithSprite(
	target *ebiten.Image,
	text string,
	x, y float64,
	scaleX, scaleY float64,
) {
	posX := x
	for _, char := range text {
		index := t.getCharIndex(char)
		if index == -1 {
			posX += t.frameWidth * scaleX
			continue
		}

		options := &ganim8.DrawOptions{
			X:      posX,
			Y:      y,
			ScaleX: scaleX,
			ScaleY: scaleY,
		}

		t.fontSprite.Draw(target, index, options)
		posX += t.frameWidth * scaleX
	}
}

func (t *TextRenderer) getCharIndex(char rune) int {
	index := -1
	switch {
	case char >= 'A' && char <= 'Z':
		index = int(char - 'A') // Uppercase letters start from index 0
	case char >= '0' && char <= '9':
		index = 26 + int(char-'0') // Numbers follow uppercase, starting at index 26
	case char >= 'a' && char <= 'z':
		index = 36 + int(char-'a') // Lowercase letters follow numbers, starting at index 36
	case char == ':':
		index = 62 // Assuming ':' is at index 62 in the sheet
	case char == '.':
		index = 63
	default:
		return -1 // Invalid or unmapped character
	}
	return index
}

func (t *TextRenderer) ScaledHeight(scale float64) float64 {
	return t.frameHeight * scale
}

func (t *TextRenderer) ScaledWidth(text string, scale float64) float64 {
	// Calculate the total width of the text based on the number of characters and scale
	return float64(len(text)) * t.frameWidth * scale
}
