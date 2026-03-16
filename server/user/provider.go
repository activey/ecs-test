package user

import (
	"ecs-test/server/user/gorm"
	"ecs-test/server/user/service"
	"go.uber.org/dig"
)

func ProvideModuleComponents(container *dig.Container) error {
	var err error
	err = container.Provide(gorm.NewUserRepository)
	if err != nil {
		return err
	}

	err = container.Provide(service.NewUserService)
	if err != nil {
		return err
	}

	err = container.Provide(service.NewAuthenticationService)
	return err
}
