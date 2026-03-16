package world

import (
	"ecs-test/server/character/domain"
	"fmt"
)

type ParticipantPredicate func(p *Participant) bool

type ParticipantId string

func (p ParticipantId) String() string {
	return string(p)
}

type Participant struct {
	Id        ParticipantId
	Character *domain.Character
	Position  Coordinates

	MovementDistance   uint
	WaitingForNextTurn bool
	NextTurnCounter    uint
}

func NewParticipant(character *domain.Character, id ParticipantId) (*Participant, error) {
	return &Participant{
		Id:               id,
		Character:        character,
		MovementDistance: character.Race.Speed.Value,
	}, nil
}

func (p *Participant) Update(timestamp int64) {
	p.updateMovementPoints()
}

func (p *Participant) RollInitiative() {
	//p.Initiative = uint(rand.Intn(20)) + p.Character.AbilityScores.Dexterity()
}

func (p *Participant) ReceiveDamage(damagePoints uint) {
	// TODO
	p.Character.ReceiveDamage(damagePoints)
}

func (p *Participant) IsAlive() bool {
	return p.Character.IsAlive()
}

func (p *Participant) CommitMovementDistance(movementPointsCommitted uint) {
	if p.MovementDistance == 0 {
		return
	}
	p.MovementDistance -= movementPointsCommitted
}

func (p *Participant) WaitForNextTurn() {
	if p.WaitingForNextTurn {
		return
	}

	p.WaitingForNextTurn = true
	p.NextTurnCounter = 6 // 6 seconds per round in realtime
}

func (p *Participant) updateMovementPoints() {
	if !p.WaitingForNextTurn {
		return
	}

	fmt.Printf("Waiting for next round: %d\n", p.NextTurnCounter)

	p.NextTurnCounter -= 1
	if p.NextTurnCounter == 0 {
		p.WaitingForNextTurn = false
		p.MovementDistance = p.Character.Race.Speed.Value
		return
	}
}
