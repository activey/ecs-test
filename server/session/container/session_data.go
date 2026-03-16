package container

import (
	"ecs-test/server/character/domain"
	"ecs-test/shared/session"
)

type SessionData struct {
	UserName  string
	SessionId session.SessionId
	Character *domain.Character
}

func NewSessionData(id session.SessionId, userName string) SessionData {
	return SessionData{
		SessionId: id,
		UserName:  userName,
	}
}

func (s SessionData) WithCharacterData(character *domain.Character) SessionData {
	if character == nil {

		return s
	}
	s.Character = character
	return s
}
