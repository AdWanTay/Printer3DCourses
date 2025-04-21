package services

import (
	"Printer3DCourses/internal/dto"
	"Printer3DCourses/internal/repositories"
	"context"
)

type CourseService struct {
	repo     repositories.CourseRepository
	userRepo repositories.UserRepository
}

func NewCourseService(repo repositories.CourseRepository, userRepo repositories.UserRepository) *CourseService {
	return &CourseService{repo: repo, userRepo: userRepo}
}

func (cr *CourseService) GetCoursesForResponse(c context.Context) (*dto.CoursesPageResponse, error) {
	courses, err := cr.repo.GetAllCourses(c)

	coursesResponse := dto.NewCoursesPageResponse(courses)
	if err != nil {
		return nil, err
	}

	return coursesResponse, nil
}

func (cr *CourseService) GetAllPaidCoursesForResponse(c context.Context, userId uint) (*dto.ProfilePageResponse, error) {
	courses, err := cr.repo.GetAllPaidCoursesByUserId(c, userId)

	user, err := cr.userRepo.GetUserById(c, userId)

	userCoursesResponse := dto.NewProfilePageResponse(courses, user)

	if err != nil {
		return nil, err
	}

	return userCoursesResponse, nil
}
