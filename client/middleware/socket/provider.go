package socket

import (
	"ecs-test/client/middleware/socket/client"
	"ecs-test/client/middleware/socket/service"
	"go.uber.org/dig"
)

func ProvideModuleComponents(container *dig.Container) error {
	err := container.Provide(service.NewSocketService)
	if err != nil {
		return err
	}

	return container.Provide(client.NewServerSocketClient)
}
