package shaders

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

var (
	//go:embed shutter.go
	shutterShaderCode []byte
	//go:embed rain.go
	rainShaderCode []byte

	ShutterShader *ebiten.Shader
	RainShader    *ebiten.Shader
)

func MustLoadShaders() {
	shader, err := ebiten.NewShader(shutterShaderCode)
	if err != nil {
		log.Panic(err)
	}
	ShutterShader = shader
}
