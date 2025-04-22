package tests

import "Printer3DCourses/internal/models"

type Test struct {
	ID        uint   `gorm:"primaryKey"`
	TestTitle string `json:"test_title"`

	CourseID uint          `gorm:"not null;OnDelete:CASCADE;"`
	Course   models.Course `gorm:"foreignKey:CourseID;references:ID"`
}
