package player

import (
	"ecs-test/client/middleware/player/service"
	"go.uber.org/dig"
)

func ProvideModuleComponents(container *dig.Container) error {
	err := container.Provide(service.NewPlayerService)
	if err != nil {
		return err
	}

	return nil
}
