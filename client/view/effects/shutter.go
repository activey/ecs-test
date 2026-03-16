package effects

import (
	"ecs-test/assets/shaders"
	"github.com/hajimehoshi/ebiten/v2"
)

type ShutterMode int

const (
	ShutterExpand ShutterMode = iota
	ShutterShrink
)

type Shutter struct {
	screenWidth, screenHeight int
	time                      float64
	active                    bool
	mode                      ShutterMode
	pixelationFactor          float64
	shutterSpeed              float64
	callback                  func()

	offscreenImage *ebiten.Image
}

func NewShutter(
	screenWidth, screenHeight int,
	pixelationFactor, shutterSpeed float64,
) *Shutter {
	return &Shutter{
		screenWidth:      screenWidth,
		screenHeight:     screenHeight,
		active:           false,
		time:             0.0,
		pixelationFactor: pixelationFactor,
		shutterSpeed:     shutterSpeed,
		offscreenImage:   ebiten.NewImage(screenWidth, screenHeight),
	}
}

func (s *Shutter) Draw(screen *ebiten.Image) {
	if !s.active {
		return
	}

	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()

	// Set shader uniforms
	op := &ebiten.DrawRectShaderOptions{}
	op.Uniforms = map[string]any{
		"Time":             float32(s.time),
		"ScreenSize":       []float32{float32(s.screenWidth), float32(s.screenHeight)},
		"Mode":             float32(s.mode),
		"PixelationFactor": float32(s.pixelationFactor),
		"ShutterSpeed":     float32(s.shutterSpeed), // Add this line
	}

	// Pass the offscreen image to the shader
	op.Images[0] = s.offscreenImage

	// Draw the shader to the offscreen image
	screen.DrawRectShader(w, h, shaders.ShutterShader, op)
}

func (s *Shutter) Update() {
	if !s.active {
		return
	}

	dt := 1.0 / float64(ebiten.TPS()) * s.shutterSpeed
	s.time += dt

	// Clamp time to 1.0
	if s.time >= 1.0 {
		s.time = 1.0
		s.active = false
		if s.callback != nil {
			s.callback()
		}
	}
}

func (s *Shutter) Start(mode ShutterMode, callback func()) {
	s.mode = mode
	s.active = true
	s.time = 0.0
	s.callback = callback
}
