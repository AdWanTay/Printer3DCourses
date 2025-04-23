package services

import (
	"Printer3DCourses/internal/dto"
	"Printer3DCourses/internal/repositories"
	"context"
)

type CourseService struct {
	repo     repositories.CourseRepository
	userRepo repositories.UserRepository
	testRepo repositories.TestRepository
}

func NewCourseService(repo repositories.CourseRepository, userRepo repositories.UserRepository, testRepo repositories.TestRepository) *CourseService {
	return &CourseService{
		repo:     repo,
		userRepo: userRepo,
		testRepo: testRepo,
	}
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
	if err != nil {
		return nil, err
	}
	progresses, err := cr.repo.GetCoursesProgress(c, userId)
	if err != nil {
		return nil, err
	}
	userCoursesResponse := dto.NewProfilePageResponse(courses, user, progresses)

	return userCoursesResponse, nil
}

func (cr *CourseService) GetCourseDetailInfo(c context.Context, userId, courseId uint) (*dto.CourseViewPageResponse, error) {
	course, err := cr.repo.GetCourseById(c, courseId)
	courseTests, err := cr.testRepo.GetTestsByCourseId(courseId)
	courseProgress, err := cr.testRepo.GetCourseProgressForUser(courseId, userId)
	testsProgresses, err := cr.testRepo.GetTestProgressesForUser(courseId, userId)
	if err != nil {
		return nil, err
	}
	return dto.NewCourseViewPageResponse(course, courseProgress, courseTests, testsProgresses), nil

}
