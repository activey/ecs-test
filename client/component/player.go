package component

import (
	"ecs-test/client/component/world"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/ganim8/v2"
	"math"
)

type PlayerDirection int

const (
	PlayerDirectionUp PlayerDirection = iota
	PlayerDirectionDown
	PlayerDirectionLeft
	PlayerDirectionRight
)

type PlayerData struct {
	currentAnimation      *ganim8.Animation
	currentDirection      PlayerDirection
	isMoving              bool
	currentNavigationPath *world.NavigationPath

	WalkingRightAnimation *ganim8.Animation
	WalkingLeftAnimation  *ganim8.Animation
	WalkingDownAnimation  *ganim8.Animation
	WalkingUpAnimation    *ganim8.Animation
	IdleUpAnimation       *ganim8.Animation
	IdleDownAnimation     *ganim8.Animation
	IdleLeftAnimation     *ganim8.Animation
	IdleRightAnimation    *ganim8.Animation
	DustSprite            *ganim8.Sprite

	VerticalShift float64
	MovementSpeed float64
}

func (d *PlayerData) DrawCurrentAnimation(screen *ebiten.Image, drawOpts *ganim8.DrawOptions) {
	w, h := d.currentAnimation.Size()
	d.currentAnimation.Draw(screen, ganim8.DrawOpts(
		drawOpts.X-(float64(w)/2*drawOpts.ScaleX),
		drawOpts.Y-(float64(h)/2*drawOpts.ScaleY)+d.VerticalShift*drawOpts.ScaleY,
		0,
		drawOpts.ScaleX,
		drawOpts.ScaleY,
	))
}

func (d *PlayerData) GoIdle() {
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

func (d *PlayerData) Update(playerTransform *transform.TransformData, tileWidth, tileHeight int) {
	d.currentAnimation.Update()

	if d.currentNavigationPath == nil {
		return
	}

	if d.currentNavigationPath.HasNodes() {
		targetNode := d.currentNavigationPath.LastNode()
		playerX, playerY := playerTransform.LocalPosition.X, playerTransform.LocalPosition.Y
		targetX, targetY := float64(targetNode.X*tileWidth)+float64(tileWidth/2), float64(targetNode.Y*tileHeight)+float64(tileHeight/2)
		direction := d.getDirection(playerX, playerY, targetX, targetY)

		d.Move(direction, playerTransform, func(newX, newY float64) bool {
			return true
		})
		if math.Abs(playerX-targetX) < 5.0 && math.Abs(playerY-targetY) < 5.0 {
			d.currentNavigationPath.RemoveLastNode()
		}
	} else {
		d.GoIdle()
		d.cancelNavigationPath()
	}
}

func (d *PlayerData) getDirection(playerX, playerY, targetX, targetY float64) PlayerDirection {
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

func (d *PlayerData) Move(
	direction PlayerDirection,
	playerTransform *transform.TransformData,
	coordinateCheck func(newX, newY float64) bool,
) {
	switch direction {
	case PlayerDirectionUp:
		if d.currentDirection != PlayerDirectionUp || !d.isMoving {
			d.currentAnimation = d.WalkingUpAnimation
		}

		newY := playerTransform.LocalPosition.Y - d.MovementSpeed/1.5
		if coordinateCheck(playerTransform.LocalPosition.X, newY) {
			playerTransform.LocalPosition.Y = newY
		}
	case PlayerDirectionDown:
		if d.currentDirection != PlayerDirectionDown || !d.isMoving {
			d.currentAnimation = d.WalkingDownAnimation
		}

		newY := playerTransform.LocalPosition.Y + d.MovementSpeed/1.5
		if coordinateCheck(playerTransform.LocalPosition.X, newY) {
			playerTransform.LocalPosition.Y = newY
		}
	case PlayerDirectionLeft:
		if d.currentDirection != PlayerDirectionLeft || !d.isMoving {
			d.currentAnimation = d.WalkingLeftAnimation
		}

		newX := playerTransform.LocalPosition.X - d.MovementSpeed
		if coordinateCheck(newX, playerTransform.LocalPosition.Y) {
			playerTransform.LocalPosition.X = newX
		}
	case PlayerDirectionRight:
		if d.currentDirection != PlayerDirectionRight || !d.isMoving {
			d.currentAnimation = d.WalkingRightAnimation
		}

		newX := playerTransform.LocalPosition.X + d.MovementSpeed
		if coordinateCheck(newX, playerTransform.LocalPosition.Y) {
			playerTransform.LocalPosition.X = newX
		}
	}

	d.currentDirection = direction
	d.isMoving = true
}

func (d *PlayerData) cancelNavigationPath() {
	if d.currentNavigationPath != nil {
		d.currentNavigationPath.Cancel()
		d.currentNavigationPath = nil
	}
}

func (d *PlayerData) MoveUp(playerTransform *transform.TransformData, coordinateCheck func(newX, newY float64) bool) {
	d.cancelNavigationPath()
	d.Move(PlayerDirectionUp, playerTransform, coordinateCheck)
}

func (d *PlayerData) MoveDown(playerTransform *transform.TransformData, coordinateCheck func(newX, newY float64) bool) {
	d.cancelNavigationPath()
	d.Move(PlayerDirectionDown, playerTransform, coordinateCheck)
}

func (d *PlayerData) MoveLeft(playerTransform *transform.TransformData, coordinateCheck func(newX, newY float64) bool) {
	d.cancelNavigationPath()
	d.Move(PlayerDirectionLeft, playerTransform, coordinateCheck)
}

func (d *PlayerData) MoveRight(playerTransform *transform.TransformData, coordinateCheck func(newX, newY float64) bool) {
	d.cancelNavigationPath()
	d.Move(PlayerDirectionRight, playerTransform, coordinateCheck)
}

func (d *PlayerData) FollowPath(path *world.NavigationPath) {
	d.currentNavigationPath = path
	fmt.Println("following path!")
}

func (d *PlayerData) IsMoving() bool {
	return d.isMoving
}

func (d *PlayerData) FollowingPath() bool {
	return d.currentNavigationPath != nil
}

func (d *PlayerData) LookDirection() PlayerDirection {
	return d.currentDirection
}

var Player = donburi.NewComponentType[PlayerData]()
var PlayerQuery = donburi.NewQuery(filter.Contains(Player, transform.Transform))
