package model

import (
	"ecs-test/shared/rules/item"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemInstanceId string

type ItemEntity struct {
	gorm.Model
	InstanceId        ItemInstanceId `gorm:"column:instance_id;unique;not null"`
	CharacterEntityId uint           `gorm:"column:character_id"`
	ItemIndex         item.Index
}

func (item *ItemEntity) BeforeCreate(tx *gorm.DB) (err error) {
	item.InstanceId = ItemInstanceId(uuid.New().String())
	return nil
}

func (ItemEntity) TableName() string {
	return "items"
}
