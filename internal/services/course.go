package services

import (
	"Printer3DCourses/internal/dto"
	"Printer3DCourses/internal/repositories"
	"context"
)

type CourseService struct {
	repo repositories.CourseRepository
}

func NewCourseService(repo repositories.CourseRepository) *CourseService {
	return &CourseService{repo: repo}
}

func (cr *CourseService) GetCoursesForResponse(c context.Context) (*dto.CoursesPageResponse, error) {
	courses, err := cr.repo.GetAllCourses(c)

	coursesResponse := dto.NewCoursesPageResponse(courses)
	if err != nil {
		return nil, err
	}

	return coursesResponse, nil
}
