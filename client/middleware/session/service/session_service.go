package service

import (
	"ecs-test/client/middleware/session/infra/rest"
	"ecs-test/shared/session"
	"fmt"
)

type SessionService struct {
	apiClient *rest.SessionApiClient
}

func NewSessionService(apiClient *rest.SessionApiClient) *SessionService {
	return &SessionService{
		apiClient: apiClient,
	}
}

type CreateSessionStatus int

const (
	CreateSessionSuccess CreateSessionStatus = iota
	CreateSessionUnauthorized
	CreateSessionUnknown
)

type CreateSessionResult struct {
	Status    CreateSessionStatus
	SessionId session.SessionId
}

func (r CreateSessionResult) Failed() bool {
	return r.Status != CreateSessionSuccess
}

func NewUnauthorizedErrorResult() CreateSessionResult {
	return CreateSessionResult{
		Status: CreateSessionUnauthorized,
	}
}

func NewUnknownErrorResult() CreateSessionResult {
	return CreateSessionResult{
		Status: CreateSessionUnknown,
	}
}

func NewSuccessResult(sessionId session.SessionId) CreateSessionResult {
	return CreateSessionResult{
		Status:    CreateSessionSuccess,
		SessionId: sessionId,
	}
}

func (s *SessionService) CreateSessionForCredentials(email string, password string) chan CreateSessionResult {
	fmt.Printf("about to send data: %s %s\n", email, password)

	resulChan := make(chan CreateSessionResult)

	go func() {
		sessionId, statusCode, err := s.apiClient.CreateSessionForCredentials(email, password)
		if err != nil {
			resulChan <- NewUnknownErrorResult()
			return
		}
		if statusCode >= 400 {
			if statusCode < 500 {
				resulChan <- NewUnauthorizedErrorResult()
			} else {
				resulChan <- NewUnknownErrorResult()
			}
			return
		}
		resulChan <- NewSuccessResult(sessionId)
	}()

	return resulChan
}
