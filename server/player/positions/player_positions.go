package positions

import (
	"ecs-test/shared/player"
	"ecs-test/shared/session"
	"ecs-test/shared/socket/payload"
)

type PlayerPositions struct {
	positionChangeChan chan *payload.PlayerMovementBroadcast

	positionMap map[session.SessionId]player.Position
}

func NewPlayerPositions() *PlayerPositions {
	return &PlayerPositions{
		positionChangeChan: make(chan *payload.PlayerMovementBroadcast, 20),

		positionMap: make(map[session.SessionId]player.Position),
	}
}

func (i *PlayerPositions) HandlerMovePayload(movePayload *payload.PlayerMovementBroadcast) {
	i.positionChangeChan <- movePayload
}

func (i *PlayerPositions) UpdatePosition(id session.SessionId, position player.Position) {
	i.positionChangeChan <- payload.NewPlayerMovementBroadcast().WithSessionId(id).WithMovementData(position.X, position.Y, 0)
}

func (i *PlayerPositions) UpdatePositionsLoop() {
	for {
		select {
		case playerMoveMessage := <-i.positionChangeChan:
			i.processMoveMessage(playerMoveMessage)
		}
	}
}

func (i *PlayerPositions) processMoveMessage(movementPayload *payload.PlayerMovementBroadcast) {
	i.positionMap[movementPayload.SessionId] = player.Position{
		X: movementPayload.X,
		Y: movementPayload.Y,
	}
}

func (i *PlayerPositions) PlayerPosition(id session.SessionId) (position player.Position, found bool) {
	position, found = i.positionMap[id]
	return
}
