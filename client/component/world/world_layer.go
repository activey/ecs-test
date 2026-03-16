package world

import (
	"ecs-test/client/camera"
	"sync"
)

type LayerID int

func (w LayerID) Id() int {
	return int(w)
}

const (
	Ground LayerID = iota
	Elevation
	Decorations
	Walkable
	OtherDecorations
)

type Layer struct {
	Tiles [][]Tile
}

func (w *Layer) ForEachTileInViewportOptimized(viewport camera.Viewport, action func(tile Tile)) {
	tileWidth := 16.0
	tileHeight := 16.0

	// Precompute viewport boundaries in tile units
	startX := int(viewport.Min.X / tileWidth)
	startY := int(viewport.Min.Y / tileHeight)
	endX := int(viewport.Max.X / tileWidth)
	endY := int(viewport.Max.Y / tileHeight)

	// Clamp to valid range
	if startX < 0 {
		startX = 0
	}
	if startY < 0 {
		startY = 0
	}
	if endX >= len(w.Tiles[0]) {
		endX = len(w.Tiles[0]) - 1
	}
	if endY >= len(w.Tiles) {
		endY = len(w.Tiles) - 1
	}

	// Iterate only over the visible tiles
	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x++ {
			tile := w.Tiles[y][x]
			if tile != nil {
				action(tile)
			}
		}
	}
}

func (w *Layer) ForEachTileInViewport(viewport camera.Viewport, action func(tile Tile)) {
	tileWidth := 16.0
	tileHeight := 16.0

	for y := range w.Tiles {
		for x := range w.Tiles[y] {
			tile := w.Tiles[y][x]
			if tile == nil {
				continue
			}

			tileX := float64(x) * tileWidth
			tileY := float64(y) * tileHeight

			if tileX+tileWidth >= viewport.Min.X && tileX <= viewport.Max.X &&
				tileY+tileHeight >= viewport.Min.Y && tileY <= viewport.Max.Y {
				action(tile)
			}
		}
	}
}

func (w *Layer) ForEachTile(action func(tile Tile)) {
	for y := 0; y < len(w.Tiles); y++ {
		for _, tile := range w.Tiles[y] {
			if tile != nil {
				action(tile)
			}
		}
	}
}

func (w *Layer) ForEachTileParallel(action func(tile Tile), numWorkers int) {
	var wg sync.WaitGroup

	rowsPerWorker := len(w.Tiles) / numWorkers
	for i := 0; i < numWorkers; i++ {
		startRow := i * rowsPerWorker
		endRow := startRow + rowsPerWorker

		if i == numWorkers-1 {
			endRow = len(w.Tiles)
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()

			for y := start; y < end; y++ {
				for _, tile := range w.Tiles[y] {
					if tile != nil {
						action(tile)
					}
				}
			}
		}(startRow, endRow)
	}
	wg.Wait()
}
