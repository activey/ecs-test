package system

import (
	"ecs-test/assets/sprites/animations"
	"ecs-test/client/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"time"
)

type PlayerDustTrace struct {
	cameraQuery *donburi.Query

	playerEntry *donburi.Entry
	camera      *donburi.Entry

	dustParticles []*component.DustParticle
	frameCounter  int
}

func NewPlayerDustTrace() *PlayerDustTrace {
	return &PlayerDustTrace{
		cameraQuery: donburi.NewQuery(filter.Contains(component.Camera, transform.Transform)),
	}
}

func (t *PlayerDustTrace) Update(e *ecs.ECS) {
	t.findPlayer(e)
	t.findCamera(e)

	if t.camera == nil || t.playerEntry == nil {
		return
	}

	playerData := component.Player.Get(t.playerEntry)
	playerTransform := transform.Transform.Get(t.playerEntry)

	if playerData.IsMoving() {
		t.EmitDustParticle(playerTransform)
		t.EmitDustParticle(playerTransform)
		//t.EmitDustParticle(playerTransform)
	}

	t.UpdateDustParticles(e.Time.DeltaTime())
}

func (t *PlayerDustTrace) EmitDustParticle(playerTransform *transform.TransformData) {
	if t.frameCounter%15 == 0 {
		dust := component.NewDustParticle(
			playerTransform.LocalPosition.X,
			playerTransform.LocalPosition.Y,
			animations.NewDustAnimation(),
		)
		t.dustParticles = append(t.dustParticles, dust)
	}
	t.frameCounter++
}

func (t *PlayerDustTrace) UpdateDustParticles(deltaTime time.Duration) {
	playerData := component.Player.Get(t.playerEntry)

	var activeParticles []*component.DustParticle
	for _, particle := range t.dustParticles {
		particle.Update(deltaTime, playerData.LookDirection())
		if particle.IsAlive() {
			activeParticles = append(activeParticles, particle)
		}
	}
	t.dustParticles = activeParticles
}

func (t *PlayerDustTrace) Draw(e *ecs.ECS, screen *ebiten.Image) {
	if t.camera == nil || t.playerEntry == nil {
		return
	}

	cameraTransform := transform.Transform.Get(t.camera)
	playerData := component.Player.Get(t.playerEntry)

	t.drawDustParticles(screen, cameraTransform, playerData.LookDirection())
}

func (t *PlayerDustTrace) drawDustParticles(
	screen *ebiten.Image,
	cameraTransform *transform.TransformData,
	direction component.PlayerDirection,
) {
	for _, particle := range t.dustParticles {
		particle.Draw(screen, cameraTransform, direction)
	}
}

func (t *PlayerDustTrace) findPlayer(e *ecs.ECS) {
	if t.playerEntry == nil {
		playerEntry, entryFound := component.PlayerQuery.First(e.World)
		if !entryFound {
			//log.Fatalf("Player entry not found!")
			return
		}
		t.playerEntry = playerEntry
	}
}

func (t *PlayerDustTrace) findCamera(e *ecs.ECS) {
	if t.camera == nil {
		cameraEntity, cameraFound := t.cameraQuery.First(e.World)
		if !cameraFound {
			// Handle case where the world map entity is not found
		}
		t.camera = cameraEntity
	}
}
