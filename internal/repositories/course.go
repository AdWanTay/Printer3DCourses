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
	// Сначала получаем все тесты
	var tests []models.Test
	err := cr.db.WithContext(c).Find(&tests).Error
	if err != nil {
		return nil, err
	}

	// Строим мапу: courseId -> общее количество тестов
	totalTestsPerCourse := make(map[uint]int64)
	for _, test := range tests {
		totalTestsPerCourse[test.CourseID]++
	}

	// Получаем все результаты тестов пользователя
	var passedScores []models.TestScore
	err = cr.db.WithContext(c).
		Preload("Test.Course").
		Where("user_id = ?", userId).
		Find(&passedScores).Error
	if err != nil {
		return nil, err
	}

	// Считаем сумму очков для каждого курса
	totalScorePerCourse := make(map[uint]int64)
	for _, score := range passedScores {
		totalScorePerCourse[score.Test.CourseID] += int64(score.Score * 100)
	}

	// Собираем финальный результат
	result := make(map[uint]int)
	for courseId, totalTests := range totalTestsPerCourse {
		if totalTests == 0 {
			result[courseId] = 0
			continue
		}
		totalScore := totalScorePerCourse[courseId]
		progress := totalScore / totalTests
		result[courseId] = int(progress)
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
