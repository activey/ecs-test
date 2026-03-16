package archetype

import (
	"ecs-test/assets/sprites/animations"
	"ecs-test/client/component"
	"ecs-test/shared/session"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func NewRemotePlayer(
	w donburi.World,
	sessionId session.SessionId,
	startPosition math.Vec2,
) *donburi.Entry {
	remotePlayer := w.Entry(
		w.Create(
			transform.Transform,
			component.RemotePlayer,
		),
	)

	p := component.RemotePlayer.Get(remotePlayer)
	p.SessionId = sessionId
	p.IdleDownAnimation = animations.NewIdleDownAnimation()
	p.IdleLeftAnimation = animations.NewIdleLeftAnimation()
	p.IdleRightAnimation = animations.NewIdleRightAnimation()
	p.IdleUpAnimation = animations.NewIdleUpAnimation()
	p.WalkingDownAnimation = animations.NewWalkingDownAnimation()
	p.WalkingUpAnimation = animations.NewWalkingUpAnimation()
	p.WalkingLeftAnimation = animations.NewWalkingLeftAnimation()
	p.WalkingRightAnimation = animations.NewWalkingRightAnimation()
	p.MovementSpeed = 1.5
	p.VerticalShift = -20 // center sprite correctly
	p.GoIdle()

	transform.Transform.Get(remotePlayer).LocalPosition = startPosition
	return remotePlayer
}
