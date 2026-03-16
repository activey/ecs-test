package archetype

import (
	"ecs-test/assets/sprites/ui"
	"ecs-test/client/component"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func NewCursor(w donburi.World, initialPosition dmath.Vec2) *donburi.Entry {
	cursor := w.Entry(
		w.Create(
			component.Cursor,
			transform.Transform,
		),
	)

	transform.Transform.Get(cursor).LocalPosition = initialPosition
	component.Cursor.Get(cursor).Sprite = ui.CursorSprite
	return cursor
}
