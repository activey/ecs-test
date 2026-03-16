package scene

import (
	"ecs-test/client/config"
	"ecs-test/client/event"
	"github.com/charmbracelet/log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type SceneFactory interface {
	Load() Scene
}

type SceneSwitcher struct {
	world         donburi.World
	width, height int

	scenes       map[string]SceneFactory
	currentScene Scene
}

func NewSceneSwitcher(world donburi.World, config config.GameClientConfig) *SceneSwitcher {
	switcher := &SceneSwitcher{
		world:  world,
		width:  config.Width,
		height: config.Height,
		scenes: make(map[string]SceneFactory),
	}

	event.SceneChangeEvent.Subscribe(world, switcher.listenSceneChange)

	return switcher
}

func (s *SceneSwitcher) AddScene(sceneName string, sceneFactory SceneFactory) {
	log.Debugf("Adding sceneFactory: %s", sceneName)
	s.scenes[sceneName] = sceneFactory
	if s.currentScene == nil {
		s.currentScene = sceneFactory.Load()
	}
}

func (s *SceneSwitcher) Switch(sceneName string) {
	scene, ok := s.scenes[sceneName]
	if !ok {
		log.Fatalf("unable to find scene: %s", sceneName)
		return
	}
	s.currentScene = scene.Load()
}

func (s *SceneSwitcher) Update() error {
	if s.currentScene == nil {
		return nil
	}
	return s.currentScene.Update()
}

func (s *SceneSwitcher) Draw(screen *ebiten.Image) {
	if s.currentScene == nil {
		return
	}
	s.currentScene.Draw(screen)
}

func (s *SceneSwitcher) listenSceneChange(w donburi.World, change event.SceneChange) {
	s.Switch(change.SceneName)
}
