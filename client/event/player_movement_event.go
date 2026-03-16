package event

import (
	"ecs-test/client/component"
	"github.com/yohamta/donburi/features/events"
)

type PlayerMovement struct {
	LookDirection component.PlayerDirection
	X, Y          float64
}

func NewPlayerMovement(direction component.PlayerDirection, x float64, y float64) PlayerMovement {
	return PlayerMovement{
		LookDirection: direction,
		X:             x,
		Y:             y,
	}
}

var PlayerMovementEvent = events.NewEventType[PlayerMovement]()
