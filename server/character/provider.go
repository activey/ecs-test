package character

import (
	"ecs-test/server/character/gorm"
	"ecs-test/server/character/service"
	"go.uber.org/dig"
)

func ProvideModuleComponents(container *dig.Container) error {
	var err error
	err = container.Provide(gorm.NewCharacterRepository)
	if err != nil {
		return err
	}

	err = container.Provide(service.NewCharacterService)
	if err != nil {
		return err
	}

	return err
}
