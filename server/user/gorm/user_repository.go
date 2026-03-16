package gorm

import (
	"context"
	"ecs-test/server/user/gorm/model"
	"ecs-test/server/user/repository"
	"errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (u UserRepository) FetchById(ctx context.Context, id uint) (*model.UserEntity, error) {
	user := &model.UserEntity{}
	if err := u.db.WithContext(ctx).First(user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (u UserRepository) FetchByUsername(ctx context.Context, username string) (*model.UserEntity, error) {
	user := &model.UserEntity{}
	if err := u.db.WithContext(ctx).Where("username = ?", username).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return UserRepository{
		db: db,
	}
}
