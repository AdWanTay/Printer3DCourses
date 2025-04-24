package repositories

import (
	"Printer3DCourses/internal/models"
	"context"
	"gorm.io/gorm"
)

type CourseRepository interface {
	GetAllCourses(c context.Context) (*[]models.Course, error)
	GetAllPaidCoursesByUserId(c context.Context, userId uint) (*[]models.Course, error)
	GetCourseById(c context.Context, id uint) (*models.Course, error)
	GetCoursesProgress(c context.Context, userId uint) (map[uint]int, error)
}
type courseRepository struct {
	db *gorm.DB
}

func (cr *courseRepository) GetCourseById(c context.Context, id uint) (*models.Course, error) {
	var course *models.Course
	err := cr.db.WithContext(c).First(&course, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (cr *courseRepository) GetCoursesProgress(c context.Context, userId uint) (map[uint]int, error) {
	var testScores []models.TestScore
	err := cr.db.WithContext(c).
		Preload("Test").
		Where("user_id = ?", userId).
		Find(&testScores).Error
	if err != nil {
		return nil, err
	}

	type progressData struct {
		total float32
		count int
	}
	courseProgress := make(map[uint]progressData)

	for _, score := range testScores {
		courseId := score.Test.CourseID
		data := courseProgress[courseId]
		data.total += score.Score
		data.count++
		courseProgress[courseId] = data
	}

	result := make(map[uint]int)
	for courseId, data := range courseProgress {
		if data.count == 0 {
			result[courseId] = 0
		} else {
			result[courseId] = int(data.total / float32(data.count) * 100)
		}
	}

	return result, nil
}

func (cr *courseRepository) GetAllPaidCoursesByUserId(c context.Context, userId uint) (*[]models.Course, error) {
	var usersCourses *[]models.UsersCourse
	err := cr.db.WithContext(c).Preload("Course").Where("user_id = ?", userId).Find(&usersCourses).Error
	if err != nil {
		return nil, err
	}
	var courses []models.Course
	for _, course := range *usersCourses {
		courses = append(courses, course.Course)
	}

	return &courses, nil
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (cr *courseRepository) GetAllCourses(c context.Context) (*[]models.Course, error) {
	var courses []models.Course
	err := cr.db.WithContext(c).Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return &courses, nil

}
