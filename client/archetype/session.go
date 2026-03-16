package archetype

import (
	"ecs-test/client/component"
	"github.com/yohamta/donburi"
)

func NewSession(w donburi.World) *donburi.Entry {
	session := w.Entry(
		w.Create(
			component.Session,
		),
	)

	return session
}
