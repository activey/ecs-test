package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strconv"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

const (
	tileSize  = 16
	gridWidth = 9
)

func lerpColor(c1, c2 color.RGBA, t float64) color.RGBA {
	return color.RGBA{
		R: uint8(float64(c1.R)*(1-t) + float64(c2.R)*t),
		G: uint8(float64(c1.G)*(1-t) + float64(c2.G)*t),
		B: uint8(float64(c1.B)*(1-t) + float64(c2.B)*t),
		A: 255,
	}
}

func drawLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{255, 0, 255, 255} // Fuchsia text color
	point := fixed.Point26_6{fixed.I(x), fixed.I(y)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

func main() {
	black := color.RGBA{0, 0, 0, 125}
	white := color.RGBA{255, 255, 255, 125}

	width := gridWidth * tileSize
	height := tileSize

	// Create an RGBA image
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.Point{}, draw.Src)

	for i := -4; i <= 4; i++ {
		t := float64(i+4) / 8.0 // Normalize between 0 and 1
		tileColor := lerpColor(black, white, t)

		x := (i + 4) * tileSize

		// Create a tile with the interpolated color
		tile := image.NewRGBA(image.Rect(0, 0, tileSize, tileSize))
		draw.Draw(tile, tile.Bounds(), &image.Uniform{tileColor}, image.Point{}, draw.Src)

		// Draw the tile onto the main image
		draw.Draw(img, image.Rect(x, 0, x+tileSize, tileSize), tile, image.Point{}, draw.Src)

		// Add the number in the center of the tile
		label := strconv.Itoa(i)
		textX := x + (tileSize / 2) - (len(label) * 3)
		textY := (tileSize / 2) + 4
		drawLabel(img, textX, textY, label)
	}

	// Save to a PNG file
	file, err := os.Create("output.png")
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatalf("failed to encode PNG: %v", err)
	}

	log.Println("Grid image saved as output.png")
}
