package domain

import (
	"ecs-test/shared/rules/class"
	"ecs-test/shared/rules/item"
	"ecs-test/shared/rules/race"
)

type Character struct {
	UserId    uint
	Name      string
	Race      *race.Race
	Class     *class.Class
	Equipment Equipment

	AbilityScores AbilityScores
	HitPoints     uint
	Level         uint
}

func (c *Character) ReceiveDamage(damagePoints uint) {
	if damagePoints >= c.HitPoints {
		c.HitPoints = 0
	} else {
		c.HitPoints -= damagePoints
	}
}

func (c *Character) IsAlive() bool {
	return c.HitPoints > 0
}

func (c *Character) Collect(item item.Item) {
	c.Equipment.CollectItem(item)
}

func NewCharacter(
	userId uint,
	name string,
	race *race.Race,
	class *class.Class,
	abilityScores AbilityScores,
	hitPoints uint,
) *Character {
	return &Character{
		UserId:        userId,
		Level:         1,
		Name:          name,
		Race:          race,
		Class:         class,
		AbilityScores: abilityScores,
		HitPoints:     hitPoints,
		Equipment:     NewEquipment(),
	}
}
