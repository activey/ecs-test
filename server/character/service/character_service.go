package service

import (
	"context"
	"ecs-test/server/character/domain"
	model2 "ecs-test/server/character/gorm/model"
	"ecs-test/server/character/repository"
	"ecs-test/server/user/service"
	"ecs-test/shared/rules/ability"
	"errors"
	"github.com/charmbracelet/log"
	"time"
)

type CharacterService struct {
	characterRepository repository.CharacterRepository
	userService         *service.Service
}

func NewCharacterService(
	repository repository.CharacterRepository,
	userService *service.Service,
) *CharacterService {
	return &CharacterService{
		characterRepository: repository,
		userService:         userService,
	}
}

func (c CharacterService) FindById(id uint) *domain.Character {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	characterEntity, err := c.characterRepository.FetchByUserId(ctx, id)
	if err != nil {
		log.Error(err)
		return nil
	}

	if characterEntity == nil {
		return nil
	}
	return characterEntity.ToCharacter()
}

func (c CharacterService) FindForUser(username string) *domain.Character {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	characterEntity, err := c.characterRepository.FetchByUsername(ctx, username)
	if err != nil {
		log.Error(err)
		return nil
	}

	if characterEntity == nil {
		return nil
	}
	return characterEntity.ToCharacter()
}

func (c CharacterService) CreateCharacter(input CreateCharacterInput) (*domain.Character, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	u := c.userService.FindById(input.UserId, ctx)
	if u == nil {
		return nil, errors.New("unable to find user")
	}

	char, err := c.characterRepository.SaveCharacter(ctx, &model2.CharacterEntity{
		UserID:     u.Id,
		Name:       input.CharacterName,
		ClassIndex: input.ClassIndex,
		RaceIndex:  input.RaceIndex,
		AbilityScores: model2.AbilityScoresEntity{
			Strength:     input.AbilityScores[ability.Strength],
			Dexterity:    input.AbilityScores[ability.Dexterity],
			Constitution: input.AbilityScores[ability.Constitution],
			Intelligence: input.AbilityScores[ability.Intelligence],
			Wisdom:       input.AbilityScores[ability.Wisdom],
			Charisma:     input.AbilityScores[ability.Charisma],
		},
	})
	if err != nil {
		return nil, err
	}
	return char.ToCharacter(), nil
}
