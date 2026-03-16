package system

import (
	"ecs-test/client/view/effects"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
)

type ShutterSystem struct {
	shutter *effects.Shutter
}

func NewShutterSystem(
	screenWidth, screenHeight int,
	pixelationFactor, shutterSpeed float64,
) *ShutterSystem {
	return &ShutterSystem{
		shutter: effects.NewShutter(screenWidth, screenHeight, pixelationFactor, shutterSpeed),
	}
}

func (s *ShutterSystem) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	s.shutter.Draw(screen)
}

func (s *ShutterSystem) Update(e *ecs.ECS) {
	s.shutter.Update()
}

func (s *ShutterSystem) Start(mode effects.ShutterMode, callback func()) {
	s.shutter.Start(mode, callback)
}
