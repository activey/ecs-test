package model

import (
	"ecs-test/server/character/domain"
	"ecs-test/server/item/model"
	"ecs-test/shared/rules"
	"ecs-test/shared/rules/class"
	"ecs-test/shared/rules/race"
	"gorm.io/gorm"
)

type CharacterEntity struct {
	gorm.Model
	Name   string
	UserID uint `gorm:"column:user_id"`

	RaceIndex     race.Index
	ClassIndex    class.Index
	AbilityScores AbilityScoresEntity
	Items         []model.ItemEntity
}

func (c CharacterEntity) ToCharacter() *domain.Character {
	char := domain.NewCharacter(
		c.UserID,
		c.Name,
		rules.AllRaces.ByIndex(c.RaceIndex),
		rules.AllClasses.ByIndex(c.ClassIndex),
		c.AbilityScores.ToAbilityScores(),
		100,
	)
	for _, item := range c.Items {
		char.Collect(rules.AllItems.ByIndex(item.ItemIndex))
	}
	return char
}

func (CharacterEntity) TableName() string {
	return "characters"
}
