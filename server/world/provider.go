package world

import (
	"ecs-test/assets"
	"github.com/charmbracelet/log"
	"github.com/lafriks/go-tiled"
	"go.uber.org/dig"
)

func ProvideModuleComponents(container *dig.Container) error {
	var err error

	err = container.Provide(provideWorldConfig)
	if err != nil {
		return err
	}

	err = container.Provide(provideTiledMap)
	if err != nil {
		return err
	}

	err = container.Provide(NewDummyConsumer)
	if err != nil {
		return err
	}

	err = container.Provide(NewWorld)
	if err != nil {
		return err
	}

	err = container.Provide(NewWorldMap)
	if err != nil {
		return err
	}

	err = container.Invoke(loadWorldMap)
	if err != nil {
		return err
	}
	return err
}

func provideWorldConfig() *Configuration {
	return NewWorldConfiguration("scenes/world2.tmx")
}

func provideTiledMap(worldConfig *Configuration) *tiled.Map {
	log.Infof("Loading world from file: %s\n", worldConfig.MapFile)
	worldMap, err := tiled.LoadFile(worldConfig.MapFile, tiled.WithFileSystem(assets.ScenesAssets))
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return worldMap
}

func loadWorldMap(p *Map, t *tiled.Map) {
	p.LoadFromTiled(t)
}
