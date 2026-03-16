package archetype

import (
	"ecs-test/assets/sprites/animations"
	"ecs-test/client/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func NewPlayer(
	w donburi.World,
	startPosition math.Vec2,
	speed float64,
) *donburi.Entry {
	player := w.Entry(
		w.Create(
			transform.Transform,
			component.Player,
		),
	)

	p := component.Player.Get(player)
	p.MovementSpeed = speed
	p.IdleDownAnimation = animations.NewIdleDownAnimation()
	p.IdleLeftAnimation = animations.NewIdleLeftAnimation()
	p.IdleRightAnimation = animations.NewIdleRightAnimation()
	p.IdleUpAnimation = animations.NewIdleUpAnimation()
	p.WalkingDownAnimation = animations.NewWalkingDownAnimation()
	p.WalkingUpAnimation = animations.NewWalkingUpAnimation()
	p.WalkingLeftAnimation = animations.NewWalkingLeftAnimation()
	p.WalkingRightAnimation = animations.NewWalkingRightAnimation()
	p.DustSprite = animations.DustSprite
	p.VerticalShift = -20 // center sprite correctly
	p.GoIdle()

	transform.Transform.Get(player).LocalPosition = startPosition
	return player
}
