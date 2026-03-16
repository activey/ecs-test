package service

import (
	"ecs-test/client/event"
	"ecs-test/client/middleware/broadcast/client"
	"ecs-test/shared/session"
	"ecs-test/shared/socket"
	"ecs-test/shared/socket/payload"
	"github.com/charmbracelet/log"
)

type BroadcastService struct {
	broadcastClient *client.ServerBroadcastClient
}

func NewBroadcastService(socketClient *client.ServerBroadcastClient) *BroadcastService {
	return &BroadcastService{broadcastClient: socketClient}
}

func (s *BroadcastService) Connect() error {
	return s.broadcastClient.Connect()
}

func (s *BroadcastService) Start() {
	s.broadcastClient.Start()
}

func (s *BroadcastService) BroadcastPlayerMovement(sessionId session.SessionId, movement event.PlayerMovement) error {
	return s.broadcastClient.SendBroadcast(
		payload.NewPlayerMovementBroadcast().
			WithSessionId(sessionId).
			WithMovementData(movement.X, movement.Y, int(movement.LookDirection)),
	)
}

func (s *BroadcastService) OnPlayerMovement(handler func(broadcast *payload.PlayerMovementBroadcast)) {
	s.broadcastClient.OnBroadcast(socket.PlayerMovementBroadcast, func(broadcast socket.Message) {
		movementBroadcast := payload.NewPlayerMovementBroadcast()
		err := broadcast.DecodePayload(movementBroadcast)
		if err != nil {
			log.Error(err)
			return
		}
		handler(movementBroadcast)
	})
}

func (s *BroadcastService) Disconnect() {
	s.broadcastClient.Disconnect()
}
