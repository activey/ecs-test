package model

import (
	"ecs-test/server/character/domain"
	"ecs-test/shared/rules/ability"
	"gorm.io/gorm"
)

type AbilityScoresEntity struct {
	gorm.Model

	CharacterEntityId uint `gorm:"column:character_id"`

	Strength,
	Dexterity,
	Constitution,
	Intelligence,
	Wisdom,
	Charisma uint
}

func (a AbilityScoresEntity) ToAbilityScores() domain.AbilityScores {
	return domain.AbilityScores{
		ability.Strength:     a.Strength,
		ability.Dexterity:    a.Dexterity,
		ability.Constitution: a.Constitution,
		ability.Intelligence: a.Intelligence,
		ability.Wisdom:       a.Wisdom,
		ability.Charisma:     a.Charisma,
	}
}

func (AbilityScoresEntity) TableName() string {
	return "ability_scores"
}
