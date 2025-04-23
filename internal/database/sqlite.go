package database

import (
	"Printer3DCourses/internal/config"
	"Printer3DCourses/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetConnection(cfg config.DatabaseConfig) (*gorm.DB, error) {
	var err error

	once.Do(func() {
		var dsn string

		switch cfg.Driver {
		case "sqlite":
			dsn = cfg.Name
			db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		case "postgres":
			dsn = fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
			)
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		default:
			err = fmt.Errorf("unsupported database driver: %s", cfg.Driver)
		}

		if err != nil {
			return
		}

		// Автоматическая миграция таблиц
		err = db.AutoMigrate(&models.User{}, &models.Course{}, &models.UsersCourse{}, &models.Test{}, &models.Question{}, &models.Answer{}, &models.TestScore{})
		err = populateDB(db)
	})
	//err = db.AutoMigrate(&models.User{}, &models.Course{}, &models.UsersCourse{})

	if db == nil {
		return nil, err
	}

	return db, nil
}

func populateDB(db *gorm.DB) error {
	authorName := "Полное Имя Автора"
	authorTgUsername := "kizzzzzaaaaa"
	courses := []models.Course{
		{
			CourseTitle:          "Основы 3D-печати: старт для новичков",
			ShortDescription:     "Описание 1",
			FullDescription:      "Полное описание",
			NumberOfParticipants: "150",
			Duration:             40,
			LatexPath:            "course1",
			Price:                2900,
			AuthorName:           authorName,
			AuthorTgUsername:     authorTgUsername,
		},
		{
			CourseTitle:          "3D-моделирование для печати в Blender",
			ShortDescription:     "Описание 2",
			FullDescription:      "Полное описание",
			NumberOfParticipants: "423",
			Duration:             35,
			LatexPath:            "course2",
			Price:                3200,
			AuthorName:           authorName,
			AuthorTgUsername:     authorTgUsername,
		},
		{
			CourseTitle:          "Печать прототипов: от модели до изделия",
			ShortDescription:     "Описание 3",
			FullDescription:      "Полное описание",
			NumberOfParticipants: "93",
			Duration:             50,
			LatexPath:            "course3",
			Price:                3500,
			AuthorName:           authorName,
			AuthorTgUsername:     authorTgUsername,
		},
		{
			CourseTitle:          "Настройка и калибровка 3D-принтера",
			ShortDescription:     "Описание 4",
			FullDescription:      "Полное описание",
			NumberOfParticipants: "45",
			Duration:             30,
			LatexPath:            "course4",
			Price:                2700,
			AuthorName:           authorName,
			AuthorTgUsername:     authorTgUsername,
		},
		{
			CourseTitle:          "Продвинутый курс по SLA и FDM технологиям",
			ShortDescription:     "Описание 5",
			FullDescription:      "Полное описание",
			NumberOfParticipants: "892",
			Duration:             45,
			LatexPath:            "course5",
			Price:                4100,
			AuthorName:           authorName,
			AuthorTgUsername:     authorTgUsername,
		},
	}
	result := db.Create(&courses)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
