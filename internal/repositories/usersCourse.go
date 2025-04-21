package repositories

import (
	"Printer3DCourses/internal/models"
	"context"
	"gorm.io/gorm"
)

type UsersCourseRepository interface {
	GetAllCoursesByUserId(c context.Context, userId int64) ([]*models.Course, error)
}

type usersCourseRepository struct {
	db *gorm.DB
}

func (u usersCourseRepository) GetAllCoursesByUserId(c context.Context, userId int64) ([]*models.Course, error) {
	var usersCourses []*models.UsersCourse
	err := u.db.WithContext(c).Where("user_id = ?", userId).Find(&usersCourses).Error
	if err != nil {
		return nil, err
	}

	//var courses []*models.Course
	//usersCourses[0].Course
	return nil, nil
}

func NewUsersCourseRepository(db *gorm.DB) UsersCourseRepository {
	return &usersCourseRepository{db: db}
}
