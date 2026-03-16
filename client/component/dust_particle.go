package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/ganim8/v2"
	"math/rand"
	"time"
)

type DustParticle struct {
	X, Y        float64
	LifeTime    time.Duration
	ElapsedTime time.Duration
	Animation   *ganim8.Animation

	dead bool
}

func NewDustParticle(x, y float64, animation *ganim8.Animation) *DustParticle {
	return &DustParticle{
		X:         x,
		Y:         y - 8,
		LifeTime:  time.Millisecond * 200,
		Animation: animation,
	}
}

func (d *DustParticle) Update(deltaTime time.Duration, direction PlayerDirection) {
	if direction == PlayerDirectionRight {
		d.X += 1
		d.Y += float64(randomInt(-10, 10)) * 0.1
	} else if direction == PlayerDirectionLeft {
		d.X -= 1
		d.Y += float64(randomInt(-10, 10)) * 0.1
	} else {
		d.X += float64(randomInt(-5, 5)) * 0.1
		if direction == PlayerDirectionDown {
			d.Y += float64(randomInt(0, 15)) * 0.1
		} else {
			d.Y -= float64(randomInt(0, 15)) * 0.1
		}
	}
	d.ElapsedTime += deltaTime
	d.Animation.Update()

	if d.Animation.IsEnd() {
		d.dead = true
	}
}

func (d *DustParticle) Draw(
	screen *ebiten.Image,
	cameraTransform *transform.TransformData,
	direction PlayerDirection,
) {

	horizontalShift := 0.0
	verticalShift := 0.0

	switch direction {
	case PlayerDirectionLeft:
		horizontalShift = 1
	case PlayerDirectionRight:
		horizontalShift = -10
	case PlayerDirectionUp:
		horizontalShift = -4
		verticalShift = 2
	case PlayerDirectionDown:
		verticalShift = -20
		horizontalShift = -4
	}

	horizontalShift *= cameraTransform.LocalScale.X
	verticalShift *= cameraTransform.LocalScale.Y

	if d.ElapsedTime < d.LifeTime {
		screenX := (d.X-cameraTransform.LocalPosition.X)*cameraTransform.LocalScale.X + horizontalShift
		screenY := (d.Y-cameraTransform.LocalPosition.Y)*cameraTransform.LocalScale.Y + verticalShift
		opts := ganim8.DrawOpts(screenX, screenY, 0, cameraTransform.LocalScale.X, cameraTransform.LocalScale.Y)

		d.Animation.Draw(screen, opts)
	}
}
func (d *DustParticle) IsAlive() bool {
	return d.dead || d.ElapsedTime < d.LifeTime
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
