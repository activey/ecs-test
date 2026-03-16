package scene

import "github.com/yohamta/donburi/ecs"

type LayerIndex struct {
	ecs.LayerID
}

func (l *LayerIndex) Next() ecs.LayerID {
	l.LayerID = l.LayerID + 1
	return l.LayerID
}
