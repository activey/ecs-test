package component

import (
	"ecs-test/shared/session"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	math2 "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/ganim8/v2"
	"math"
	"time"
)

type PositionUpdate struct {
	X, Y float64
	Time time.Time
}

type RemotePlayerData struct {
	currentAnimation *ganim8.Animation
	currentDirection PlayerDirection
	isMoving         bool
	latestUpdate     time.Time

	WalkingRightAnimation *ganim8.Animation
	WalkingLeftAnimation  *ganim8.Animation
	WalkingDownAnimation  *ganim8.Animation
	WalkingUpAnimation    *ganim8.Animation
	IdleUpAnimation       *ganim8.Animation
	IdleDownAnimation     *ganim8.Animation
	IdleLeftAnimation     *ganim8.Animation
	IdleRightAnimation    *ganim8.Animation

	VerticalShift float64
	MovementSpeed float64
	SessionId     session.SessionId
}

func (d *RemotePlayerData) DrawCurrentAnimation(screen *ebiten.Image, drawOpts *ganim8.DrawOptions) {
	w, h := d.currentAnimation.Size()

	opts := ganim8.DrawOpts(
		drawOpts.X-(float64(w)/2*drawOpts.ScaleX),
		drawOpts.Y-(float64(h)/2*drawOpts.ScaleY)+d.VerticalShift*drawOpts.ScaleY,
		0,
		drawOpts.ScaleX,
		drawOpts.ScaleY,
	)
	opts.ColorM.Scale(0.5, 1, 0.5, 0.9)

	d.currentAnimation.Draw(screen, opts)
}

func (d *RemotePlayerData) GoIdle() {
	if !d.isMoving && d.currentAnimation != nil {
		return
	}
	switch d.currentDirection {
	case PlayerDirectionUp:
		d.currentAnimation = d.IdleUpAnimation
	case PlayerDirectionDown:
		d.currentAnimation = d.IdleDownAnimation
	case PlayerDirectionLeft:
		d.currentAnimation = d.IdleLeftAnimation
	case PlayerDirectionRight:
		d.currentAnimation = d.IdleRightAnimation
	}
	d.isMoving = false
}

func (d *RemotePlayerData) Update(playerTransform *transform.TransformData, deltaTime float64) {
	d.currentAnimation.Update()
	d.GoIdle()
}

func (d *RemotePlayerData) getDirection(playerX, playerY, targetX, targetY float64) PlayerDirection {
	if math.Abs(playerX-targetX) > math.Abs(playerY-targetY) {
		if playerX < targetX {
			return PlayerDirectionRight
		} else {
			return PlayerDirectionLeft
		}
	} else {
		if playerY < targetY {
			return PlayerDirectionDown
		} else {
			return PlayerDirectionUp
		}
	}
}

func (d *RemotePlayerData) Move(
	direction PlayerDirection,
	playerTransform *transform.TransformData,
	x float64, y float64,
) {
	switch direction {
	case PlayerDirectionUp:
		if d.currentDirection != PlayerDirectionUp || !d.isMoving {
			d.currentAnimation = d.WalkingUpAnimation
		}
	case PlayerDirectionDown:
		if d.currentDirection != PlayerDirectionDown || !d.isMoving {
			d.currentAnimation = d.WalkingDownAnimation
		}
	case PlayerDirectionLeft:
		if d.currentDirection != PlayerDirectionLeft || !d.isMoving {
			d.currentAnimation = d.WalkingLeftAnimation
		}
	case PlayerDirectionRight:
		if d.currentDirection != PlayerDirectionRight || !d.isMoving {
			d.currentAnimation = d.WalkingRightAnimation
		}
	}

	if playerTransform.LocalPosition.X != x || playerTransform.LocalPosition.Y != y {
		playerTransform.LocalPosition = math2.NewVec2(x, y)
	}

	d.currentDirection = direction
	d.isMoving = true
}

func (d *RemotePlayerData) IsMoving() bool {
	return d.isMoving
}

func (d *RemotePlayerData) LookDirection() PlayerDirection {
	return d.currentDirection
}

func (d *RemotePlayerData) UpdateMovementTime(t time.Time) {
	d.latestUpdate = t
}

func (d *RemotePlayerData) LatestUpdateAfter(t time.Time) bool {
	return d.latestUpdate.After(t)
}

var RemotePlayer = donburi.NewComponentType[RemotePlayerData]()
var RemotePlayerQuery = donburi.NewQuery(filter.Contains(RemotePlayer, transform.Transform))
