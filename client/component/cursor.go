package component

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/ganim8/v2"
)

type CursorData struct {
	Sprite *ganim8.Sprite
}

var Cursor = donburi.NewComponentType[CursorData]()
