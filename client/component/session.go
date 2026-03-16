package component

import (
	"ecs-test/shared/session"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type SessionData struct {
	sessionId session.SessionId
}

func (d *SessionData) SessionId() session.SessionId {
	return d.sessionId
}

func (d *SessionData) SetSessionId(id session.SessionId) {
	d.sessionId = id
}

var Session = donburi.NewComponentType[SessionData]()
var SessionQuery = donburi.NewQuery(filter.Contains(Session))
