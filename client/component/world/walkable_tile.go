package world

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

type WalkableTile struct {
	BaseTile
	IsWalkable      bool
	ElevationHeight int
}

func (t WalkableTile) DebugDraw(screen *ebiten.Image, options *ebiten.DrawImageOptions, back *ebiten.Image) {
	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Translate(t.X, t.Y)
	ops.GeoM.Concat(options.GeoM)

	screen.DrawImage(back, ops)

	// Calculate the correct text size using font.BoundString
	bounds, _ := font.BoundString(basicfont.Face7x13, fmt.Sprintf("%d", t.ElevationHeight))
	textWidth := (bounds.Max.X - bounds.Min.X).Ceil()  // Calculate width
	textHeight := (bounds.Max.Y - bounds.Min.Y).Ceil() // Calculate height

	// Center the text inside the tile
	textX := float64(t.Width)/2 - float64(textWidth)/2
	textY := float64(t.Height)/2 + float64(textHeight)/2

	// Set up the options for text drawing
	textOps := &ebiten.DrawImageOptions{}
	textOps.GeoM.Translate(t.X+textX, t.Y+textY)
	textOps.GeoM.Concat(options.GeoM)

	// Draw the elevation height text (centered) using DrawWithOptions
	text.DrawWithOptions(
		screen,
		fmt.Sprintf("%d", t.ElevationHeight),
		basicfont.Face7x13,
		textOps,
	)
}
