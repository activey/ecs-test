package session

import (
	"ecs-test/server/infra/http"
	"ecs-test/server/session/container"
	"ecs-test/server/session/infra/rest"
	"go.uber.org/dig"
)

func ProvideModuleComponents(c *dig.Container) error {
	var err error

	err = c.Provide(container.NewSessionContainer)
	if err != nil {
		return err
	}

	err = c.Provide(rest.NewSessionController)
	if err != nil {
		return err
	}

	// register controllers routes
	return c.Invoke(registerRoutes)
}

func registerRoutes(c *rest.SessionController, r http.RouteProvider) {
	c.RegisterRoutes(r)
}
