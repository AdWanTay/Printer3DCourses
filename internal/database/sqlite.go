package database

import (
	"Printer3DCourses/internal/config"
	"Printer3DCourses/internal/database/seeders"
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

		err = db.AutoMigrate(&models.User{}, &models.Course{}, &models.UsersCourse{}, &models.Test{}, &models.Question{}, &models.Answer{}, &models.TestScore{})
		err = populateDB(db)
		if err != nil {
			fmt.Errorf("populateDB(db)")
			return
		}
	})

	if db == nil {
		return nil, err
	}

	return db, nil
}
func populateDB(db *gorm.DB) error {
	err := seeders.SeedCourses(db)
	if err != nil {
		return err
	}
	err = seeders.SeedTestsAndQuestionsForCourse1(db)
	if err != nil {
		return err
	}
	err = seeders.SeedTestsAndQuestionsForCourse2(db)
	if err != nil {
		return err
	}
	err = seeders.SeedTestsAndQuestionsForCourse3(db)
	if err != nil {
		return err
	}
	err = seeders.SeedTestsAndQuestionsForCourse4(db)
	if err != nil {
		return err
	}
	err = seeders.SeedTestsAndQuestionsForCourse5(db)
	if err != nil {
		return err
	}
	fmt.Println("Database seeded successfully")
	return nil
}
