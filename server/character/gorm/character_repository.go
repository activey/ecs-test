package gorm

import (
	"context"
	"ecs-test/server/character/gorm/model"
	"ecs-test/server/character/repository"
	"errors"
	"gorm.io/gorm"
)

type CharacterRepository struct {
	db *gorm.DB
}

func (c CharacterRepository) SaveCharacter(ctx context.Context, character *model.CharacterEntity) (*model.CharacterEntity, error) {
	result := c.db.WithContext(ctx).Create(character)
	if result.Error != nil {
		return nil, result.Error
	}
	return character, nil
}

func (c CharacterRepository) FetchByUserId(ctx context.Context, id uint) (*model.CharacterEntity, error) {
	character := &model.CharacterEntity{}

	if err := c.db.WithContext(ctx).
		Joins("JOIN users ON users.id = characters.user_id").
		Preload("Items").
		Preload("AbilityScores").
		Where("users.id = ?", id).
		First(character).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return character, nil
}

func (c CharacterRepository) FetchByUsername(ctx context.Context, username string) (*model.CharacterEntity, error) {
	character := &model.CharacterEntity{}

	if err := c.db.WithContext(ctx).
		Joins("JOIN users ON users.id = characters.user_id").
		Preload("Items").
		Preload("AbilityScores").
		Where("users.username = ?", username).
		First(character).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return character, nil
}

func NewCharacterRepository(db *gorm.DB) repository.CharacterRepository {
	return CharacterRepository{
		db: db,
	}
}
