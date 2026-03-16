package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"image/color"
	"log"
)

type WalkableLayerReader struct {
}

func (r WalkableLayerReader) LoadWalkableLayer(worldMap *tiled.Map, walkableId LayerID) *WalkableLayer {
	tmxLayer := worldMap.Layers[walkableId.Id()]
	if tmxLayer == nil {
		log.Panic("Walkable layer not found!")
	}

	debugBack := ebiten.NewImage(worldMap.TileWidth, worldMap.TileHeight)
	debugBack.Fill(color.RGBA{R: 128, G: 128, B: 128, A: 128}) // Only create and fill once

	walkableLayer := NewWalkableLayer(worldMap)
	for y, layerTile := range tmxLayer.Tiles {
		x := y % worldMap.Width
		row := y / worldMap.Width

		if walkableLayer.Tiles[row] == nil {
			walkableLayer.Tiles[row] = make([]*WalkableTile, worldMap.Width)
		}

		if layerTile != nil && !layerTile.IsNil() {
			tileSet := layerTile.Tileset
			tile, err := tileSet.GetTilesetTile(layerTile.ID)
			if err != nil {
				continue
			}

			heightProp := tile.Properties.GetInt("height")
			walkableLayer.Tiles[row][x] = &WalkableTile{
				BaseTile: BaseTile{
					X:      float64(x * tileSet.TileWidth),
					Y:      float64(row * tileSet.TileHeight),
					Width:  tileSet.TileWidth,
					Height: tileSet.TileHeight,
				},

				IsWalkable:      true,
				ElevationHeight: heightProp,
			}
		}
	}
	walkableLayer.BuildNavigationGrid()
	return walkableLayer
}

func NewWalkableLayerReader() *WalkableLayerReader {
	return &WalkableLayerReader{}
}
