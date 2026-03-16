package item

import "ecs-test/shared/rules/ability"

type DamageType string

func (d DamageType) String() string {
	return string(d)
}

const (
	Bludgeoning DamageType = "Bludgeoning"
	Piercing    DamageType = "Piercing"
	Slashing    DamageType = "Slashing"
	Cold        DamageType = "Cold"
	Fire        DamageType = "Fire"
	Lightning   DamageType = "Lightning"
	Thunder     DamageType = "Thunder"
	Acid        DamageType = "Acid"
	Poison      DamageType = "Poison"
	Radiant     DamageType = "Radiant"
	Necrotic    DamageType = "Necrotic"
	Force       DamageType = "Force"
	Psychic     DamageType = "Psychic"
)

type Damage struct {
	Type         DamageType
	Name         string
	PlusModifier *ability.Score
	Dice         string
}
