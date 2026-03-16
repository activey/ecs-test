package service

import (
	"ecs-test/client/middleware/socket/client"
	"ecs-test/shared/session"
	"ecs-test/shared/socket"
	"ecs-test/shared/socket/payload"
	"github.com/charmbracelet/log"
)

type SocketService struct {
	socketClient *client.ServerSocketClient
}

func NewSocketService(socketClient *client.ServerSocketClient) *SocketService {
	return &SocketService{socketClient: socketClient}
}

func (s *SocketService) Connect() error {
	return s.socketClient.Connect()
}

func (s *SocketService) RequestJoin(sessionId session.SessionId) (joinResultChan chan *payload.JoinResponse, e error) {
	joinResultChan = make(chan *payload.JoinResponse)

	request := payload.NewPlayerJoinRequest(sessionId)
	responseChan, err := s.socketClient.SendRequest(request)
	if err != nil {
		e = err
		log.Error(err)
		return
	}

	go func() {
		select {
		case responseData := <-responseChan:
			response := payload.NewPlayerJoinResponsePayload()
			err := response.DecodeFrom(responseData)
			if err != nil {
				e = err
				log.Error(err)
				return
			}
			joinResultChan <- response
		}
	}()

	return
}

func (s *SocketService) OnPlayerJoined(handler func(broadcast *payload.JoinBroadcast)) {
	s.socketClient.OnBroadcast(socket.PlayerJoinBroadcast, func(broadcast socket.Message) {
		joinBroadcast := payload.NewJoinBroadcast()
		err := broadcast.DecodePayload(joinBroadcast)
		if err != nil {
			log.Error(err)
			return
		}
		handler(joinBroadcast)
	})
}

func (s *SocketService) OnConnectionClosed(handlerFunc func(err error)) {
	s.socketClient.OnConnectionClosed(handlerFunc)
}

func (s *SocketService) Disconnect() error {
	return s.socketClient.Disconnect()
}
