package archetype

import (
	"ecs-test/client/component"
	"github.com/yohamta/donburi"
)

func NewDebug(w donburi.World, enabled bool) *donburi.Entry {
	debug := w.Entry(
		w.Create(
			component.Debug,
		),
	)

	debugData := component.Debug.Get(debug)
	if enabled {
		debugData.EnableDebug()
	}
	return debug
}
