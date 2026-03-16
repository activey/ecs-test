package world

import (
	"ecs-test/client/camera"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/lafriks/go-tiled"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/ganim8/v2"
	"image/color"
)

type WalkableLayer struct {
	Tiles                 [][]*WalkableTile
	debugTileBackground   *ebiten.Image
	TileWidth, TileHeight int
	navigationGrid        *NavigationGrid
}

func NewWalkableLayer(worldMap *tiled.Map) *WalkableLayer {
	debugBack := ebiten.NewImage(worldMap.TileWidth, worldMap.TileHeight)
	debugBack.Fill(color.RGBA{R: 128, G: 128, B: 128, A: 128}) // Only create and fill once

	return &WalkableLayer{
		Tiles:               make([][]*WalkableTile, worldMap.Height),
		debugTileBackground: debugBack,
		TileWidth:           worldMap.TileWidth,
		TileHeight:          worldMap.TileHeight,
	}
}

func (w *WalkableLayer) forEachTileInViewport(viewport camera.Viewport, action func(tile *WalkableTile)) {
	for y := range w.Tiles {
		for x := range w.Tiles[y] {
			tile := w.Tiles[y][x]
			if tile == nil {
				continue
			}

			tileX := float64(x) * float64(w.TileWidth)
			tileY := float64(y) * float64(w.TileHeight)

			if tileX+float64(w.TileWidth) >= viewport.Min.X && tileX <= viewport.Max.X &&
				tileY+float64(w.TileHeight) >= viewport.Min.Y && tileY <= viewport.Max.Y {
				action(tile)
			}
		}
	}
}

func (w *WalkableLayer) forEachTile(action func(tile *WalkableTile)) {
	for _, row := range w.Tiles {
		for _, tile := range row {
			if tile == nil {
				continue
			}
			action(tile)
		}
	}
}

func (w *WalkableLayer) IsValidCoordinate(x, y float64) bool {
	for _, row := range w.Tiles {
		for _, tile := range row {
			if tile == nil {
				continue
			}
			if x >= tile.X && x < tile.X+float64(tile.Width) &&
				y >= tile.Y && y < tile.Y+float64(tile.Height) {
				return tile.IsWalkable
			}
		}
	}
	return false
}

func (w *WalkableLayer) Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions, viewport camera.Viewport) {
	w.forEachTileInViewport(viewport, func(tile *WalkableTile) {
		tile.DebugDraw(screen, options, w.debugTileBackground)
	})
}

func (w *WalkableLayer) BuildNavigationGrid() {
	fmt.Print("Building navigation grid...")
	w.navigationGrid = NewNavigationGrid()
	w.forEachTile(func(tile *WalkableTile) {
		w.navigationGrid.AddNode(tile.TileX(), tile.TileY(), true)
	})
	w.navigationGrid.BuildNeighbors()
	fmt.Println("Done!")
}

func (w *WalkableLayer) DrawNavigationPath(
	path *NavigationPath,
	color color.RGBA,
	screen *ebiten.Image,
	cameraTransform *transform.TransformData,
) {
	if len(path.nodes) < 2 {
		return
	}

	for i := 0; i < len(path.nodes)-1; i++ {
		startNode := path.nodes[i]
		endNode := path.nodes[i+1]

		if startNode == nil || endNode == nil {
			return
		}

		startX, startY := startNode.Vec2().
			Mul(dmath.NewVec2(float64(w.TileWidth), float64(w.TileHeight))).
			Add(dmath.NewVec2(float64(w.TileWidth)/2, float64(w.TileHeight)/2)). // Center the line inside the tile
			Sub(cameraTransform.LocalPosition).
			Mul(cameraTransform.LocalScale).
			XY()

		endX, endY := endNode.Vec2().
			Mul(dmath.NewVec2(float64(w.TileWidth), float64(w.TileHeight))).
			Add(dmath.NewVec2(float64(w.TileWidth)/2, float64(w.TileHeight)/2)). // Center the line inside the tile
			Sub(cameraTransform.LocalPosition).
			Mul(cameraTransform.LocalScale).
			XY()

		vector.StrokeLine(screen, float32(startX), float32(startY), float32(endX), float32(endY), 5, color, false)
	}
}

func (w *WalkableLayer) DrawTargetFlag(
	position dmath.Vec2,
	animation *ganim8.Animation,
	screen *ebiten.Image,
	cameraTransform *transform.TransformData,
) {
	spriteWidth := animation.Sprite().W()
	spriteHeight := animation.Sprite().H()

	verticalShift := 10.0

	X, Y := position.
		Mul(dmath.NewVec2(float64(w.TileWidth), float64(w.TileHeight))).
		Sub(dmath.NewVec2(float64(spriteWidth)/4, float64(spriteHeight)/2+verticalShift)).
		Sub(cameraTransform.LocalPosition).
		Mul(cameraTransform.LocalScale).
		XY()
	animation.Draw(screen, ganim8.DrawOpts(X, Y, 0, cameraTransform.LocalScale.X, cameraTransform.LocalScale.Y))
}

func (w *WalkableLayer) GetNodeFromPixel(x float64, y float64, cameraTransform *transform.TransformData) *NavigationNode {
	return w.navigationGrid.GetNodeFromPixel(
		x,
		y,
		w.TileWidth,
		w.TileHeight,
		cameraTransform,
	)
}
