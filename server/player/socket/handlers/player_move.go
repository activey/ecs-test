package handlers

import (
	"ecs-test/server/player/positions"
	"ecs-test/server/session/container"
	"ecs-test/shared/socket"
	"ecs-test/shared/socket/payload"
	"github.com/charmbracelet/log"
)

type PlayerBroadcast struct {
	sessionContainer *container.SessionContainer
	playerPositions  *positions.PlayerPositions
}

func NewPlayerBroadcast(
	sessionContainer *container.SessionContainer,
	playerPositions *positions.PlayerPositions,
) *PlayerBroadcast {
	return &PlayerBroadcast{sessionContainer: sessionContainer, playerPositions: playerPositions}
}

func (i *PlayerBroadcast) OnPlayerMoveMessage(playerMoveMessage socket.Message) *socket.Message {
	movementPayload := &payload.PlayerMovementBroadcast{}
	err := playerMoveMessage.DecodePayload(movementPayload)
	if err != nil {
		log.Error(err)
		return nil
	}

	//_, found := i.sessionContainer.GetSessionData(movementPayload.SessionId)
	//if found {
	//fmt.Println("player movement:", movementPayload.X, movementPayload.Y)
	i.playerPositions.HandlerMovePayload(movementPayload)

	movementBroadcast, err := socket.NewBroadcast(
		payload.NewPlayerMovementBroadcast().
			WithSessionId(movementPayload.SessionId).
			WithMovementData(movementPayload.X, movementPayload.Y, movementPayload.Direction),
	)
	if err == nil {
		return &movementBroadcast
	} else {
		log.Error(err)
	}
	//}

	return nil
}
