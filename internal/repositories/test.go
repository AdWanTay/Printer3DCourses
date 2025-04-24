package repositories

import (
	"Printer3DCourses/internal/models"
	"context"
	"errors"
	"gorm.io/gorm"
)

type TestRepository interface {
	GetTestById(id uint) (*models.Test, error)
	GetTestsByCourseId(courseId uint) (*[]models.Test, error)
	GetCourseProgressForUser(courseID, userID uint) (int, error)
	GetTestProgressesForUser(courseId, userId uint) (map[uint]int, error)
	GetQuestionsWithCorrectAnswers(ctx context.Context, testID uint) (*[]models.Question, error)
	SaveTestScore(c context.Context, testScore *models.TestScore) error
	GetTestScoreByUserAndTest(ctx context.Context, userID, testID uint) (*models.TestScore, error)
	UpdateTestScore(ctx context.Context, score *models.TestScore) error
}

type testRepository struct {
	db *gorm.DB
}

func (t testRepository) GetTestsByCourseId(courseId uint) (*[]models.Test, error) {
	var tests *[]models.Test
	err := t.db.Where("course_id = ?", courseId).Find(&tests).Error
	return tests, err
}

func (t testRepository) GetTestById(id uint) (*models.Test, error) {
	var tests models.Test
	err := t.db.Preload("Questions.Answers").Find(&tests, "id = ?", id).Error
	return &tests, err
}
func (t *testRepository) GetCourseProgressForUser(courseId, userId uint) (int, error) {
	var totalTests int64
	err := t.db.Model(&models.Test{}).Where("course_id = ?", courseId).Count(&totalTests).Error
	if err != nil {
		return 0, err
	}

	if totalTests == 0 {
		return 0, nil // если у курса нет тестов, прогресс — 0
	}

	var passedScores []models.TestScore
	err = t.db.
		Joins("JOIN tests ON tests.id = test_scores.test_id").
		Where("tests.course_id = ? AND test_scores.user_id = ?", courseId, userId).
		Find(&passedScores).Error
	if err != nil {
		return 0, err
	}

	var totalScore int
	for _, score := range passedScores {
		totalScore += int(score.Score * 100)
	}

	return totalScore, nil
}

func (t *testRepository) GetTestProgressesForUser(courseId, userId uint) (map[uint]int, error) {
	var tests []models.Test
	err := t.db.Where("course_id = ?", courseId).Find(&tests).Error
	if err != nil {
		return nil, err
	}

	testIDs := make([]uint, 0, len(tests))
	for _, test := range tests {
		testIDs = append(testIDs, test.ID)
	}

	var scores []models.TestScore
	err = t.db.
		Where("user_id = ? AND test_id IN ?", userId, testIDs).
		Find(&scores).Error
	if err != nil {
		return nil, err
	}

	testProgresses := make(map[uint]int)

	for _, test := range tests {
		testProgresses[test.ID] = 0
	}

	for _, score := range scores {
		testProgresses[score.TestID] = int(score.Score * 100)
	}

	return testProgresses, nil
}

func (t *testRepository) GetQuestionsWithCorrectAnswers(ctx context.Context, testID uint) (*[]models.Question, error) {
	var questions []models.Question

	err := t.db.WithContext(ctx).
		Preload("Answers", "is_correct = ?", true).
		Where("test_id = ?", testID).
		Find(&questions).Error

	if err != nil {
		return nil, err
	}

	filtered := make([]models.Question, 0)
	for _, q := range questions {
		if len(q.Answers) > 0 {
			filtered = append(filtered, q)
		}
	}
	return &filtered, nil
}

func (t *testRepository) SaveTestScore(c context.Context, testScore *models.TestScore) error {
	return t.db.WithContext(c).Create(testScore).Error
}

func (t *testRepository) GetTestScoreByUserAndTest(ctx context.Context, userID, testID uint) (*models.TestScore, error) {
	var score models.TestScore
	err := t.db.WithContext(ctx).
		Where("user_id = ? AND test_id = ?", userID, testID).
		First(&score).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // не решал
	}
	if err != nil {
		return nil, err
	}
	return &score, nil
}

func (t *testRepository) UpdateTestScore(ctx context.Context, score *models.TestScore) error {
	return t.db.WithContext(ctx).Save(score).Error
}

func NewTestRepository(db *gorm.DB) TestRepository {
	return &testRepository{db: db}
}
