package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func NewDatabaseConnection(configuration *DatabaseConfiguration) *gorm.DB {
	db, err := gorm.Open(postgres.Open(configuration.ConnectionURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}
	return db
}
