package world

import (
	"github.com/disintegration/imaging"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
	"github.com/yohamta/ganim8/v2"
	"image"
	"log"
	"time"
)

type LayerRenderer struct {
	cachedLayerTilesets map[string]*ebiten.Image
	cachedTiles         map[uint32]*ebiten.Image
	cachedAnimations    map[uint32]*ganim8.Sprite
}

func NewLayerRenderer() *LayerRenderer {
	return &LayerRenderer{
		cachedLayerTilesets: make(map[string]*ebiten.Image),
		cachedTiles:         make(map[uint32]*ebiten.Image),
		cachedAnimations:    make(map[uint32]*ganim8.Sprite),
	}
}

func (w *LayerRenderer) LoadWorldLayer(worldMap *tiled.Map, layerId LayerID) *Layer {
	tmxLayer := worldMap.Layers[layerId.Id()]
	layer := &Layer{
		Tiles: make([][]Tile, worldMap.Height),
	}

	for y, layerTile := range tmxLayer.Tiles {
		x := y % worldMap.Width
		row := y / worldMap.Width

		if layer.Tiles[row] == nil {
			layer.Tiles[row] = make([]Tile, worldMap.Width)
		}

		tileset := layerTile.Tileset
		if tileset == nil {
			continue
		}

		cachedTileSetImage := w.cachedLayerTilesets[tileset.Name]
		if cachedTileSetImage == nil {
			tileSetFile := "assets/" + tileset.GetFileFullPath(tileset.Image.Source)
			println("loading image", tileSetFile)

			loadedImage, _, err := ebitenutil.NewImageFromFile(tileSetFile)
			if err != nil {
				log.Panicf("failed to open tileset image: %v", err)
			}
			w.cachedLayerTilesets[tileset.Name] = loadedImage
			cachedTileSetImage = loadedImage
		}
		//
		if tileset.Tiles != nil {
			tile, err := tileset.GetTilesetTile(layerTile.ID)
			if err != nil {
				tileImage := w.getTileImage(layerTile)
				layer.Tiles[row][x] = &StaticTile{
					BaseTile: BaseTile{
						X:      float64(x * tileset.TileWidth),
						Y:      float64(row * tileset.TileHeight),
						Width:  tileset.TileWidth,
						Height: tileset.TileHeight,
					},
					Image: ebiten.NewImageFromImage(tileImage),
				}
				continue
			}

			if tile.Animation == nil || len(tile.Animation) == 0 {
				tileImage := w.getTileImage(layerTile)
				layer.Tiles[row][x] = &StaticTile{
					BaseTile: BaseTile{
						X:      float64(x * tileset.TileWidth),
						Y:      float64(row * tileset.TileHeight),
						Width:  tileset.TileWidth,
						Height: tileset.TileHeight,
					},
					Image: ebiten.NewImageFromImage(tileImage),
				}
			} else {
				animationSprite := w.getTileAnimation(layerTile, tile, cachedTileSetImage)
				animation := ganim8.NewAnimation(animationSprite, 100*time.Millisecond)
				layer.Tiles[row][x] = &AnimatedTile{
					BaseTile: BaseTile{
						X:      float64(x * tileset.TileWidth),
						Y:      float64(row * tileset.TileHeight),
						Width:  tileset.TileWidth,
						Height: tileset.TileHeight,
					},
					Animation: animation,
				}
			}
		} else {
			tileImage := w.getTileImage(layerTile)
			layer.Tiles[row][x] = &StaticTile{
				BaseTile: BaseTile{
					X:      float64(x * tileset.TileWidth),
					Y:      float64(row * tileset.TileHeight),
					Width:  tileset.TileWidth,
					Height: tileset.TileHeight,
				},
				Image: ebiten.NewImageFromImage(tileImage),
			}
		}
	}
	return layer
}

// Simplified image caching logic
func (w *LayerRenderer) getTileImage(layerTile *tiled.LayerTile) *ebiten.Image {
	tileSet := layerTile.Tileset
	cacheKey := tileSet.FirstGID + layerTile.ID

	if cachedTile, ok := w.cachedTiles[cacheKey]; ok {
		return cachedTile
	}

	rect := layerTile.GetTileRect()
	cachedTileSetImage := w.cachedLayerTilesets[tileSet.Name]

	cropped := imaging.Crop(cachedTileSetImage, rect)
	if layerTile.HorizontalFlip {
		cropped = imaging.FlipH(cropped)
	}
	if layerTile.VerticalFlip {
		cropped = imaging.FlipV(cropped)
	}

	tileImage := ebiten.NewImageFromImage(cropped)
	w.cachedTiles[cacheKey] = ebiten.NewImageFromImage(tileImage)

	return tileImage
}

func (w *LayerRenderer) getTileAnimation(layerTile *tiled.LayerTile, tile *tiled.TilesetTile, tileSetImage *ebiten.Image) *ganim8.Sprite {
	cachedAnimation := w.cachedAnimations[layerTile.ID]
	if cachedAnimation != nil {
		return cachedAnimation
	}

	tileSet := layerTile.Tileset
	var animationFrames []*image.Rectangle
	for _, frame := range tile.Animation {
		// Get the rectangle of each animation frame
		rect := tileSet.GetTileRect(frame.TileID)
		animationFrames = append(animationFrames, &rect)
	}
	animationSprite := ganim8.NewSprite(tileSetImage, animationFrames)
	w.cachedAnimations[tileSet.FirstGID+layerTile.ID] = animationSprite

	return animationSprite
}

func (w *LayerRenderer) Cleanup() {
	//for _, w := range w.cachedLayerTilesets {
	//	w.Deallocate()
	//}
	//for _, w := range w.cachedTiles {
	//	w.Deallocate()
	//}

	clear(w.cachedTiles)
	clear(w.cachedLayerTilesets)
	clear(w.cachedAnimations)
}
