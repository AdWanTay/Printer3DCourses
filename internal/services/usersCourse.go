package services

import (
	"Printer3DCourses/internal/repositories"
	"context"
)

type UsersCourseService struct {
	repo repositories.UsersCourseRepository
}

func NewUsersCourseService(repo repositories.UsersCourseRepository) *UsersCourseService {
	return &UsersCourseService{repo: repo}
}

func (u *UsersCourseService) BuyCourse(c context.Context, userId, courseId uint) error {
	err := u.repo.AddUserToCourse(c, userId, courseId)
	if err != nil {
		return err
	}
	return nil
}
