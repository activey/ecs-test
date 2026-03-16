package system

import (
	giff2 "ecs-test/client/view/giff"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
	"golang.org/x/image/draw"
	"image"
)

type GifRender struct {
	GIFImage         *giff2.GIFImage  // Holds GIF frames and image data
	Composite        *image.RGBA      // Composite image used to build the full frame
	CurrentEbitenImg *ebiten.Image    // Cached ebiten image for rendering
	currentIdx       int              // Current frame index
	player           *giff2.GIFPlayer // Player to control GIF playback
	needsRedraw      bool             // Flag to indicate if a new frame needs to be drawn
}

// NewGifRender loads the GIF and initializes the player
func NewGifRender(gifPath string, speed float64) (*GifRender, error) {
	gifImage, err := giff2.LoadGIFImageFromPath(gifPath)
	if err != nil {
		return nil, err
	}

	// Initialize composite image
	bounds := gifImage.Get(0).Bounds()
	composite := image.NewRGBA(bounds)

	// Create an Ebiten image once (reuse it for efficiency)
	currentEbitenImg := ebiten.NewImage(bounds.Dx(), bounds.Dy())

	// Initialize player
	player := giff2.NewPlayer(gifImage, speed)

	// Create GifRender
	br := &GifRender{
		GIFImage:         gifImage,
		Composite:        composite,
		CurrentEbitenImg: currentEbitenImg,
		currentIdx:       0,
		player:           player,
		needsRedraw:      true, // Initial frame needs to be drawn
	}

	// Register an observer for when the player triggers a frame update
	br.player.AddObserver(func() {
		br.currentIdx = (br.currentIdx + 1) % br.GIFImage.Length()
		br.needsRedraw = true
	})

	return br, nil
}

// Draw renders the current frame onto the screen
func (m *GifRender) Draw(screen *ebiten.Image) {
	if m.needsRedraw {
		// Get the current frame from the GIF
		currentFrame := m.GIFImage.Get(m.currentIdx)

		// Draw the current frame onto the composite image only when needed
		draw.Draw(m.Composite, currentFrame.Bounds(), currentFrame, image.Point{}, draw.Src)

		// Update the existing Ebiten image with the new frame using ReplacePixels (no re-creation)
		m.CurrentEbitenImg.WritePixels(m.Composite.Pix)

		m.needsRedraw = false // Reset the flag
	}

	// Render the cached Ebiten image to the screen, scaled to 800x600
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(800/float64(m.Composite.Bounds().Dx()), 600/float64(m.Composite.Bounds().Dy()))
	op.ColorScale.Scale(1, 0.8, 0.7, 0.8)
	screen.DrawImage(m.CurrentEbitenImg, op)
}

// Update advances the animation every frame
func (m *GifRender) Update(e *ecs.ECS) {
	m.player.Update()
}
