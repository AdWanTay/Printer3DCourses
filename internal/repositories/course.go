package repositories

import (
	"Printer3DCourses/internal/models"
	"context"
	"gorm.io/gorm"
)

type CourseRepository interface {
	GetAllCourses(c context.Context) (*[]models.Course, error)
}
type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (cr *courseRepository) GetAllCourses(c context.Context) (*[]models.Course, error) {
	var course []models.Course
	err := cr.db.WithContext(c).Find(&course).Error
	if err != nil {
		return nil, err
	}
	return &course, nil

}
