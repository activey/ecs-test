package socket

import (
	"ecs-test/shared/session"
)

type Connection interface {
	Write(message Message) error
	SetSessionId(id session.SessionId)
	SessionId() session.SessionId
}
