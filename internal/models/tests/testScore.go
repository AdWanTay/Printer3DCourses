package tests

import "Printer3DCourses/internal/models"

type TestScore struct {
	ID uint64 `gorm:"primaryKey"`

	TestID uint `gorm:"not null;OnDelete:CASCADE;"`
	Test   Test `gorm:"foreignKey:TestID;references:ID"`

	UserID uint        `gorm:"not null;OnDelete:CASCADE;"`
	User   models.User `gorm:"foreignKey:UserID;references:ID"`
}
