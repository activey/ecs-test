package session

import "github.com/google/uuid"

type SessionId string

func (s SessionId) String() string {
	return string(s)
}

func NewSessionId() (SessionId, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return SessionId(newUUID.String()), nil
}
