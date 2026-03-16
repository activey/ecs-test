package archetype

import (
	"ecs-test/client/component"
	"github.com/yohamta/donburi"
)

func NewPlayerLock(
	w donburi.World,
) *donburi.Entry {
	lock := w.Entry(
		w.Create(
			component.PlayerLock,
		),
	)
	return lock
}
