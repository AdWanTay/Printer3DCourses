package database

import (
	"Printer3DCourses/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetConnection() (*gorm.DB, error) {
	if db == nil {
		database, err := gorm.Open(sqlite.Open("data.sqlite"), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		// Автоматическое создание таблиц по моделям
		err = database.AutoMigrate(
			&models.User{},
		)
		if err != nil {
			return nil, err
		}
		db = database
	}
	return db, nil
}
