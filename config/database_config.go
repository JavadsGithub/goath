package config

import (
	"github.com/JavadsGithub/goath/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(AppConfig.DB_NAME), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.User{})
	return db
}
