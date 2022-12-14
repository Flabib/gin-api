package database

import (
	"gin-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Instance() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database.")
	}

	db.AutoMigrate(&models.Note{})

	return db
}
