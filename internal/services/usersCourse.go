package services

import (
	"Printer3DCourses/internal/repositories"
	"context"
	"errors"
)

type UsersCourseService struct {
	repo     repositories.UsersCourseRepository
	testRepo repositories.TestRepository
}

func NewUsersCourseService(repo repositories.UsersCourseRepository, testRepo repositories.TestRepository) *UsersCourseService {
	return &UsersCourseService{repo: repo, testRepo: testRepo}
}

func (u *UsersCourseService) BuyCourse(c context.Context, userId, courseId uint) error {
	err := u.repo.AddUserToCourse(c, userId, courseId)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersCourseService) IsBought(c context.Context, userId, id uint, param string) (bool, error) {
	var courseId uint
	switch param {
	case "course":
		courseId = id
	case "test":
		test, err := u.testRepo.GetTestById(id)
		if err != nil {
			return false, err
		}
		courseId = test.CourseID
	default:
		return false, errors.New("invalid param")
	}

	course, err := u.repo.GetCourseByUserIdAndCourseId(c, userId, courseId)
	if err != nil || course == nil {
		return false, err
	}
	return true, nil
}
