package repositories

import (
	"Printer3DCourses/internal/models"
	"context"
	"errors"
	"gorm.io/gorm"
)

type UsersCourseRepository interface {
	AddUserToCourse(ctx context.Context, userId, courseId uint) error
	GetCourseByUserIdAndCourseId(c context.Context, userId, courseId uint) (*models.UsersCourse, error)
}

type usersCourseRepository struct {
	db *gorm.DB
}

func (r *usersCourseRepository) AddUserToCourse(c context.Context, userId, courseId uint) error {
	var existing models.UsersCourse

	err := r.db.WithContext(c).
		Where("user_id = ? AND course_id = ?", userId, courseId).
		First(&existing).Error

	if err == nil {
		return errors.New("user already exists")
	}

	entry := &models.UsersCourse{
		UserID:   userId,
		CourseID: courseId,
	}
	return r.db.WithContext(c).Create(entry).Error
}

func (r *usersCourseRepository) GetCourseByUserIdAndCourseId(c context.Context, userId, courseId uint) (*models.UsersCourse, error) {
	var courses *models.UsersCourse
	err := r.db.WithContext(c).Where("user_id = ? AND course_id = ?", userId, courseId).First(&courses).Error
	return courses, err
}

func NewUsersCourseRepository(db *gorm.DB) UsersCourseRepository {
	return &usersCourseRepository{db: db}
}
