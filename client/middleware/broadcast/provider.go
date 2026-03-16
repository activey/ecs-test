package broadcast

import (
	"ecs-test/client/middleware/broadcast/client"
	"ecs-test/client/middleware/broadcast/service"
	"go.uber.org/dig"
)

func ProvideModuleComponents(container *dig.Container) error {
	err := container.Provide(service.NewBroadcastService)
	if err != nil {
		return err
	}

	return container.Provide(client.NewServerBroadcastClient)
}
