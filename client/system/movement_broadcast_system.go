package system

import (
	"ecs-test/client/component"
	"ecs-test/client/event"
	broadcastService "ecs-test/client/middleware/broadcast/service"
	"github.com/charmbracelet/log"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type MovementBroadcastSystem struct {
	broadcastService *broadcastService.BroadcastService
	initialized      bool

	world        donburi.World
	sessionQuery *donburi.Query
	sessionEntry *donburi.Entry
}

func NewMovementBroadcast(broadcastService *broadcastService.BroadcastService) *MovementBroadcastSystem {
	return &MovementBroadcastSystem{
		broadcastService: broadcastService,
		sessionQuery:     donburi.NewQuery(filter.Contains(component.Session)),
	}
}

func (t *MovementBroadcastSystem) Update(e *ecs.ECS) {
	t.findSessionComponent(e)
	if t.initialized {
		return
	}

	t.world = e.World
	event.PlayerMovementEvent.Subscribe(e.World, t.publishPlayerMovement)

	t.initialized = true
}

func (t *MovementBroadcastSystem) findSessionComponent(ecs *ecs.ECS) {
	if t.sessionEntry == nil {
		entry, ok := t.sessionQuery.First(ecs.World)
		if ok {
			t.sessionEntry = entry
		}
	}
}

func (t *MovementBroadcastSystem) publishPlayerMovement(w donburi.World, movement event.PlayerMovement) {
	session := component.Session.Get(t.sessionEntry)

	err := t.broadcastService.BroadcastPlayerMovement(session.SessionId(), movement)
	if err != nil {
		log.Error(err)
	}
}
