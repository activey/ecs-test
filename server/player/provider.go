package player

import (
	"ecs-test/server/infra/broadcast"
	socketServer "ecs-test/server/infra/socket"
	"ecs-test/server/player/positions"
	"ecs-test/server/player/socket/handlers"
	"ecs-test/shared/socket"
	"go.uber.org/dig"
)

func ProvideModuleComponents(c *dig.Container) error {
	var err error

	err = c.Provide(handlers.NewPlayerJoinHandler)
	if err != nil {
		return err
	}

	err = c.Provide(handlers.NewPlayerBroadcast)
	if err != nil {
		return err
	}

	err = c.Provide(positions.NewPlayerPositions)
	if err != nil {
		return err
	}

	// register server message handlers
	err = c.Invoke(registerMessageHandlers)
	if err != nil {
		return err
	}

	err = c.Invoke(registerBroadcastListeners)
	if err != nil {
		return err
	}

	return c.Invoke(updatePlayersPositions)
}

func registerMessageHandlers(
	s *socketServer.Server,
	pi *handlers.PlayerJoinHandler,
) {
	s.OnDisconnect(pi.OnPlayerDisconnect)
	s.OnRequest(socket.PlayerJoinRequest, pi.OnPlayerJoinMessage)
}

func registerBroadcastListeners(
	b *broadcast.Server,
	px *handlers.PlayerBroadcast,
) {
	b.OnBroadcast(socket.PlayerMovementBroadcast, px.OnPlayerMoveMessage)
}

func updatePlayersPositions(
	pp *positions.PlayerPositions,
) {
	go pp.UpdatePositionsLoop()
}
