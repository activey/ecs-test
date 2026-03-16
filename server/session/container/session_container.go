package container

import (
	"ecs-test/server/character/service"
	"ecs-test/shared/session"
	"sync"
)

type SessionContainer struct {
	sync.RWMutex

	sessions         map[session.SessionId]SessionData
	characterService *service.CharacterService
}

func NewSessionContainer(characterService *service.CharacterService) *SessionContainer {
	return &SessionContainer{
		sessions:         map[session.SessionId]SessionData{},
		characterService: characterService,
	}
}

func (c *SessionContainer) NewSessionForUser(user string) (session.SessionId, error) {
	c.Lock()
	defer c.Unlock()

	userCharacter := c.characterService.FindForUser(user)

	sessionId, err := session.NewSessionId()
	c.sessions[sessionId] = NewSessionData(sessionId, user).WithCharacterData(userCharacter)
	return sessionId, err
}

func (c *SessionContainer) FindExistingSessionForUser(user string) (SessionData, bool) {
	c.RLock()
	defer c.RUnlock()

	for _, data := range c.sessions {
		if data.UserName == user {
			return data, true
		}
	}
	return SessionData{}, false
}

func (c *SessionContainer) GetSessionData(id session.SessionId) (SessionData, bool) {
	c.RLock()
	defer c.RUnlock()

	sessionData, ok := c.sessions[id]
	return sessionData, ok
}
