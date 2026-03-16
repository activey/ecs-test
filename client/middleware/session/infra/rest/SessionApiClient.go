package rest

import (
	"ecs-test/client/config"
	"ecs-test/shared/session"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/go-resty/resty/v2"
)

type SessionApiClient struct {
	client *resty.Client
}

func NewSessionApiClient(config config.GameClientConfig) *SessionApiClient {
	client := resty.New()
	client.SetBaseURL(fmt.Sprintf("http://%s:8664", config.ServerAddress))
	return &SessionApiClient{
		client: client,
	}
}

func (s *SessionApiClient) CreateSessionForCredentials(email string, password string) (
	sessionId session.SessionId,
	statusCode int,
	err error,
) {
	s.client.SetBasicAuth(email, password)
	response, err := s.client.R().Post("/session")
	if err != nil {
		return "", 500, err
	}
	if response.IsError() {
		log.Error(response.String())
		return "", response.StatusCode(), nil
	}

	sId := session.SessionId(response.Header().Get("X-Session-ID"))
	return sId, response.StatusCode(), nil
}
