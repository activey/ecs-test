package model

import (
	"ecs-test/server/user/domain"
	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model

	Username, Password string
	//CharacterEntity    *model.CharacterEntity `gorm:"foreignKey:UserID"`
}

func (UserEntity) TableName() string {
	return "users"
}

func (e UserEntity) ToUser() *domain.User {
	return domain.NewUser(
		e.ID,
		e.Username,
		e.Password,
	)
}
