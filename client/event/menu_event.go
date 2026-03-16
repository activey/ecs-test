package event

import (
	"github.com/yohamta/donburi/features/events"
)

type MenuSelection int

const (
	JoinWorld MenuSelection = iota
	Quit
)

var MenuSelectionEvent = events.NewEventType[MenuSelection]()
