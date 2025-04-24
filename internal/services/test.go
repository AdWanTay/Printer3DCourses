package services

import (
	"Printer3DCourses/internal/dto"
	"Printer3DCourses/internal/models"
	"Printer3DCourses/internal/repositories"
	"context"
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

func (s *TestService) SaveTestResult(c context.Context, userId uint, request *dto.TestAnswersRequest) (float32, error) {
	questionsWithCorrectAnswers, err := s.repo.GetQuestionsWithCorrectAnswers(c, request.TestId)
	if err != nil {
		return 0, err
	}

	correctSum := 0
	for _, question := range *questionsWithCorrectAnswers {
		if answer, ok := request.Result[question.ID]; ok {
			if len(question.Answers) > 0 && answer == question.Answers[0].ID {
				correctSum++
			}
		}
	}
	result := float32(correctSum) / float32(len(*questionsWithCorrectAnswers))

	// Проверка: решал ли пользователь тест
	existingScore, err := s.repo.GetTestScoreByUserAndTest(c, userId, request.TestId)
	if err != nil {
		return 0, err
	}

	if existingScore != nil {
		// Если результат лучше — обновим
		if result > float32(existingScore.Score) {
			existingScore.Score = result
			err := s.repo.UpdateTestScore(c, existingScore)
			if err != nil {
				return 0, err
			}
		}
	} else {
		// Если не решал — создадим
		err = s.repo.SaveTestScore(c, &models.TestScore{
			UserID: userId,
			TestID: request.TestId,
			Score:  result,
		})
		if err != nil {
			return 0, err
		}
	}

	return result, nil
}
