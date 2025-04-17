package repositories

import (
	"Printer3DCourses/internal/models"
	"context"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(c context.Context, email string) (*models.User, error)
	CreateUser(c context.Context, user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u userRepository) GetUserByEmail(c context.Context, email string) (*models.User, error) {
	var user models.User
	err := u.db.WithContext(c).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userRepository) CreateUser(c context.Context, user *models.User) error {
	return u.db.WithContext(c).Create(user).Error
}
