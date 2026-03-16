package handlers

import (
	"ecs-test/server/player/positions"
	"ecs-test/server/session/container"
	"ecs-test/server/world"
	"ecs-test/shared/player"
	"ecs-test/shared/session"
	"ecs-test/shared/socket"
	"ecs-test/shared/socket/codec"
	"ecs-test/shared/socket/payload"
	"github.com/charmbracelet/log"
	"sync"
)

type Connections map[session.SessionId]socket.Connection

type PlayerJoinHandler struct {
	connections      Connections
	sessionContainer *container.SessionContainer
	playerPositions  *positions.PlayerPositions

	world    *world.World
	worldMap *world.Map

	codec *codec.CborCodec
}

func NewPlayerJoinHandler(
	sessionContainer *container.SessionContainer,
	playerPositions *positions.PlayerPositions,
	worldMap *world.Map,
	world *world.World,
) *PlayerJoinHandler {
	return &PlayerJoinHandler{
		world:            world,
		worldMap:         worldMap,
		sessionContainer: sessionContainer,
		playerPositions:  playerPositions,
		connections:      make(Connections),
		codec:            codec.NewCborCodec(),
	}
}

func (i *PlayerJoinHandler) OnPlayerJoinMessage(
	conn socket.Connection,
	request socket.Message,
) (response socket.Message) {
	requestPayload := &payload.JoinRequest{}
	err := request.DecodePayload(requestPayload)
	if err != nil {
		log.Error(err)
	}

	//i.world.JoinCharacter()

	//time.Sleep(3 * time.Second)

	sessionData, found := i.sessionContainer.GetSessionData(requestPayload.SessionId)
	if !found {
		log.Errorf("Unable to find session with id: %s", requestPayload.SessionId)
		response, err = socket.NewResponse(
			request.RequestID,
			socket.PlayerJoinResponse,
			payload.NewPlayerJoinResponsePayload().WithStatus(payload.JoinFailedUnknownSession),
		)
		return
	}

	_, alreadyThere := i.connections[requestPayload.SessionId]
	if alreadyThere {
		response, err = socket.NewResponse(
			request.RequestID,
			socket.PlayerJoinResponse,
			payload.NewPlayerJoinResponsePayload().WithStatus(payload.JoinFailedAlreadyThere),
		)
		return
	}

	// random location from map
	position := player.NewPosition(i.worldMap.RandomLocation())

	// broadcast player joined info to other players
	i.broadcastPlayerJoined(sessionData.SessionId, sessionData.UserName, position)

	log.Infof("Player joined with SessionID: %s\n", requestPayload.SessionId)
	conn.SetSessionId(requestPayload.SessionId)
	i.connections[requestPayload.SessionId] = conn
	i.playerPositions.UpdatePosition(requestPayload.SessionId, position)

	response, err = i.prepareJoinResponse(request, requestPayload.SessionId, position)
	if err != nil {
		log.Error(err)

		response, err = socket.NewResponse(
			request.RequestID,
			socket.PlayerJoinResponse,
			payload.NewPlayerJoinResponsePayload().WithStatus(payload.JoinFailedSystemError),
		)
	}
	return
}

func (i *PlayerJoinHandler) prepareJoinResponse(
	request socket.Message,
	sessionId session.SessionId,
	position player.Position,
) (socket.Message, error) {
	otherPlayers := make([]player.Player, 0)

	// collecting other players
	for sessId, _ := range i.connections {
		if sessionId == sessId {
			continue
		}
		data, dataExists := i.sessionContainer.GetSessionData(sessId)
		if dataExists {
			playerPosition, positionFound := i.playerPositions.PlayerPosition(data.SessionId)
			if positionFound {
				otherPlayers = append(otherPlayers, player.Player{
					Name:      data.UserName,
					SessionId: data.SessionId,
					Position:  playerPosition,
				})
			}
		}
	}

	return socket.NewResponse(
		request.RequestID,
		socket.PlayerJoinResponse,
		payload.NewPlayerJoinResponsePayload().
			WithPosition(position).
			WithOtherPlayers(otherPlayers...),
	)
}

func (i *PlayerJoinHandler) OnPlayerDisconnect(conn socket.Connection) {
	sessionId := conn.SessionId()
	log.Infof("Player with session id: [%s] disconnected, removing connection", sessionId)

	delete(i.connections, sessionId)
}

func (i *PlayerJoinHandler) broadcastPlayerJoined(
	id session.SessionId,
	name string,
	position player.Position,
) {
	broadcast, err := socket.NewBroadcast(
		payload.NewJoinBroadcast().
			WithSessionId(id).
			WithWithUserName(name).
			WithPosition(position),
	)
	if err != nil {
		log.Error(err)
		return
	}

	i.broadcast(broadcast, func(a socket.Connection) bool {
		return a.SessionId() != id
	})
}

func (i *PlayerJoinHandler) broadcast(m socket.Message, evalFunc func(a socket.Connection) bool) {
	wg := sync.WaitGroup{}

	for _, connection := range i.connections {
		if evalFunc(connection) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := connection.Write(m)
				if err != nil {
					return
				}
			}()
		}
	}
	wg.Wait()
}
