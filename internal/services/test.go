package services

import (
	"Printer3DCourses/internal/dto"
	"Printer3DCourses/internal/repositories"
)

type TestService struct {
	repo repositories.TestRepository
}

func NewTestService(repo repositories.TestRepository) *TestService {
	return &TestService{repo: repo}
}

func (s *TestService) GetTestByIdForResponse(id uint) (*dto.TestResponse, error) {
	test, err := s.repo.GetTestById(id)

	if err != nil {
		return nil, err
	}

	response := dto.NewTestResponse(test, len(test.Questions))
	return response, err
}
