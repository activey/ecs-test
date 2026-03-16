package widgets

import (
	"ecs-test/assets/sprites/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/ganim8/v2"
	"image"
)

type Panel struct {
	scale float64
	width int
}

func NewPanel(scale float64, width int) *Panel {
	return &Panel{
		scale: scale,
		width: width,
	}
}

func (t *Panel) Update(view *furex.View) {
	sprite := ui.PanelSprite // The sprite containing the yellow panel tiles
	w := sprite.W()          // The original size of each tile
	tileWidth := float64(w)

	// Calculate the scaled tile size
	scaledTileWidth := tileWidth * t.scale

	// Calculate the number of columns and rows, adjusting for scale
	numCols := int(float64(t.width) / scaledTileWidth)

	view.SetWidth(int(scaledTileWidth * float64(numCols)))
}

func (t *Panel) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	sprite := ui.PanelSprite // The sprite containing the yellow panel tiles
	w, h := sprite.Size()    // The original size of each tile
	tileWidth := float64(w)
	tileHeight := float64(h)

	// Calculate the scaled tile size
	scaledTileWidth := tileWidth * t.scale
	scaledTileHeight := tileHeight * t.scale

	// Calculate the number of columns and rows, adjusting for scale
	numCols := int(float64(t.width) / scaledTileWidth)
	numRows := int(float64(frame.Dy()) / scaledTileHeight)

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			x := float64(frame.Min.X) + float64(col)*scaledTileWidth
			y := float64(frame.Min.Y) + float64(row)*scaledTileHeight

			tileIndex := t.getTileIndex(numRows, numCols, row, col)

			sprite.Draw(screen, tileIndex, &ganim8.DrawOptions{
				X:      x,
				Y:      y,
				ScaleX: t.scale,
				ScaleY: t.scale,
			})
		}
	}
}

func (t *Panel) getTileIndex(numRows, numCols, row, col int) int {
	var tileIndex = 0

	if row == 0 {
		if col == 0 {
			tileIndex = 0
		} else if col == numCols-1 {
			tileIndex = 2
		} else {
			tileIndex = 1
		}
	} else if row == numRows-1 {
		if col == 0 {
			tileIndex = 6
		} else if col == numCols-1 {
			tileIndex = 8
		} else {
			tileIndex = 7
		}
	} else {
		if col == 0 {
			tileIndex = 3
		} else if col == numCols-1 {
			tileIndex = 5
		} else {
			tileIndex = 4
		}
	}
	return tileIndex
}
