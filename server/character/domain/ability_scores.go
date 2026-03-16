package domain

import "ecs-test/shared/rules/ability"

type AbilityScores map[ability.Score]uint

func (a AbilityScores) Strength() uint {
	return a[ability.Strength]
}

func (a AbilityScores) Dexterity() uint {
	return a[ability.Dexterity]
}

func (a AbilityScores) Constitution() uint {
	return a[ability.Constitution]
}
