package repositories

import (
	"Printer3DCourses/internal/models"
	"gorm.io/gorm"
)

type TestRepository interface {
	GetTestById(id uint) (*models.Test, error)
}

type testRepository struct {
	db *gorm.DB
}

func (t testRepository) GetTestById(id uint) (*models.Test, error) {
	var tests models.Test
	err := t.db.Preload("Questions.Answers").Find(&tests, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &tests, nil
}

func NewTestRepository(db *gorm.DB) TestRepository {
	return &testRepository{db: db}
}
