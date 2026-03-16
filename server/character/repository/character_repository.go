package repository

import (
	"context"
	"ecs-test/server/character/gorm/model"
)

type CharacterRepository interface {
	FetchByUserId(ctx context.Context, id uint) (*model.CharacterEntity, error)
	FetchByUsername(ctx context.Context, username string) (*model.CharacterEntity, error)
	SaveCharacter(ctx context.Context, character *model.CharacterEntity) (*model.CharacterEntity, error)
}
