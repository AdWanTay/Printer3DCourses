package services

import "Printer3DCourses/internal/repositories"

type TestService struct {
	repo repositories.TestRepository
}

func NewTestService(repo repositories.TestRepository) *TestService {
	return &TestService{repo: repo}
}
