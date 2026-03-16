package service

import (
	"ecs-test/client/archetype"
	"ecs-test/client/component"
	"ecs-test/shared/session"
	"github.com/charmbracelet/log"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type PlayerService struct {
	world donburi.World

	spawned bool
}

func NewPlayerService(world donburi.World) *PlayerService {
	return &PlayerService{world: world}
}

func (p *PlayerService) SpawnPlayer(x, y float64) {
	if p.spawned {
		entry, ok := component.PlayerQuery.First(p.world)
		if ok {
			t := transform.Transform.Get(entry)
			t.LocalPosition = math.NewVec2(x, y)
		}

		return
	}

	log.Infof("Spawning at: %f %f\n", x, y)
	archetype.NewPlayer(p.world, math.Vec2{X: x, Y: y}, 1.5)
	p.spawned = true
}

func (p *PlayerService) SpawnRemotePlayer(id session.SessionId, x float64, y float64) {
	archetype.NewRemotePlayer(p.world, id, math.NewVec2(x, y))
}
