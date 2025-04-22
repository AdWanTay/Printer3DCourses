package repositories

import "gorm.io/gorm"

type TestRepository interface {
}

type testRepository struct {
	db *gorm.DB
}

func NewTestRepository(db *gorm.DB) TestRepository {
	return &testRepository{db: db}
}
