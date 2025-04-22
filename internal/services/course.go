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

func (cr *CourseService) GetCoursesForResponse(c context.Context, userId *uint) (*dto.CoursesPageResponse, error) {
	courses, err := cr.repo.GetAllCourses(c)
	paidIndexes := make(map[uint]struct{})

	if userId != nil {
		paidCourses, err := cr.repo.GetAllPaidCoursesByUserId(c, *userId)
		if err != nil {
			return nil, err
		}
		for _, course := range *paidCourses {
			paidIndexes[course.ID] = struct{}{}
		}
	}

	coursesResponse := dto.NewCoursesPageResponse(courses, paidIndexes)
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
