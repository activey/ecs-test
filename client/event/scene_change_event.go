package event

import (
	"github.com/yohamta/donburi/features/events"
)

type SceneChange struct {
	SceneName string
}

func NewSceneChange(sceneName string) SceneChange {
	return SceneChange{
		SceneName: sceneName,
	}
}

var SceneChangeEvent = events.NewEventType[SceneChange]()
