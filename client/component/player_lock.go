package component

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type PlayerLockData struct {
	locked bool
}

func (l *PlayerLockData) IsLocked() bool {
	return l.locked
}

func (l *PlayerLockData) Unlock() {
	l.locked = false
}

func (l *PlayerLockData) Lock() {
	l.locked = true
}

var PlayerLock = donburi.NewComponentType[PlayerLockData]()
var PlayerLockQuery = donburi.NewQuery(filter.Contains(PlayerLock))
