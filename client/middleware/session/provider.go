package session

import (
	"ecs-test/client/middleware/session/infra/rest"
	"ecs-test/client/middleware/session/service"
	"go.uber.org/dig"
)

func ProvideModuleComponents(container *dig.Container) error {
	err := container.Provide(service.NewSessionService)
	if err != nil {
		return err
	}

	err = container.Provide(rest.NewSessionApiClient)
	if err != nil {
		return err
	}

	return nil
}
