package player

import "ecs-test/shared/session"

type Player struct {
	Name      string
	SessionId session.SessionId
	Position  Position
}
