package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
	"time"
)

type Tile interface {
	Update(time.Duration)
	Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions)
}

type BaseTile struct {
	X, Y          float64 // pixel coordinates
	Width, Height int     // in pixels
}

func (b BaseTile) TileX() int {
	return int(b.X) / b.Width
}

func (b BaseTile) TileY() int {
	return int(b.Y) / b.Height
}

// StaticTile - static tile implementation
type StaticTile struct {
	BaseTile
	Image *ebiten.Image
}

func (s StaticTile) Update(time.Duration) {}

func (s StaticTile) Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions) {
	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Translate(s.X, s.Y)
	ops.GeoM.Concat(options.GeoM)
	screen.DrawImage(s.Image, ops)
}

// AnimatedTile - animated tile implementation
type AnimatedTile struct {
	BaseTile
	Animation     *ganim8.Animation
	FrameDuration float64
}

func (t *AnimatedTile) Update(deltaTime time.Duration) {
	t.Animation.UpdateWithDelta(deltaTime)
	//t.Animation.Update()
}

func (t *AnimatedTile) Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions) {
	scaleX, scaleY := options.GeoM.Element(0, 0), options.GeoM.Element(1, 1)
	transX, transY := options.GeoM.Element(0, 2), options.GeoM.Element(1, 2)

	opts := ganim8.DrawOpts(t.X, t.Y)
	opts.SetScale(scaleX, scaleY)

	opts.X = t.X*scaleX + transX
	opts.Y = t.Y*scaleY + transY

	t.Animation.Draw(screen, opts)
}
