package service

import (
	"ecs-test/server/character/domain"
	"ecs-test/shared/rules/class"
	"ecs-test/shared/rules/race"
)

type CreateCharacterInput struct {
	UserId        uint
	CharacterName string
	RaceIndex     race.Index
	ClassIndex    class.Index
	AbilityScores domain.AbilityScores
}

func NewCreateCharacterInput(
	userId uint,
	characterName string,
	raceIndex race.Index,
	classIndex class.Index,
	abilityScores domain.AbilityScores,
) CreateCharacterInput {
	return CreateCharacterInput{
		UserId:        userId,
		CharacterName: characterName,
		RaceIndex:     raceIndex,
		ClassIndex:    classIndex,
		AbilityScores: abilityScores,
	}
}
