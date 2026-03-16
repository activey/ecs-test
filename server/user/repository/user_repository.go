package repository

import (
	"context"
	"ecs-test/server/user/gorm/model"
)

type UserRepository interface {
	FetchByUsername(ctx context.Context, username string) (*model.UserEntity, error)
	FetchById(ctx context.Context, id uint) (*model.UserEntity, error)
}
