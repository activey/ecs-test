package world

import (
	"github.com/lafriks/go-tiled"
	"log"
	"math/rand"
	"time"
)

const (
	walkableLayerId int = 3
)

type Map struct {
	tiles [][]*MapTile
}

func NewWorldMap() *Map {
	return &Map{}
}

func (m *Map) LoadFromTiled(tiledMap *tiled.Map) {
	walkableLayer := tiledMap.Layers[walkableLayerId]
	if walkableLayer == nil {
		log.Panic("Walkable layer not found!")
	}

	tiles := make([][]*MapTile, tiledMap.Height)
	for y, layerTile := range walkableLayer.Tiles {
		x := y % tiledMap.Width
		row := y / tiledMap.Width

		if tiles[row] == nil {
			tiles[row] = make([]*MapTile, tiledMap.Width)
		}

		if layerTile != nil && !layerTile.IsNil() {
			tileSet := layerTile.Tileset
			tile, err := tileSet.GetTilesetTile(layerTile.ID)
			if err != nil {
				continue
			}

			heightProp := tile.Properties.GetInt("height")
			tiles[row][x] = &MapTile{
				X:               float64(x * tileSet.TileWidth),
				Y:               float64(row * tileSet.TileHeight),
				Width:           tileSet.TileWidth,
				Height:          tileSet.TileHeight,
				ElevationHeight: heightProp,
			}
		}
	}
	m.tiles = tiles
}

func (m *Map) RandomLocation() (X, Y float64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(m.tiles) == 0 {
		return 0, 0
	}

	for {
		row := r.Intn(len(m.tiles))
		col := r.Intn(len(m.tiles[row]))

		tile := m.tiles[row][col]
		if tile != nil {
			X = tile.X + float64(tile.Width/2)
			Y = tile.Y + float64(tile.Height/2)
			return
		}
	}

}
