package services

import (
	"Printer3DCourses/internal/dto"
	"Printer3DCourses/internal/models"
	"Printer3DCourses/internal/repositories"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(c context.Context, input dto.RegistrationRequest) (*models.User, error) {
	existingUser, err := s.repo.GetUserByEmail(c, input.Email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("user with this email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	newUser := &models.User{
		LastName:    input.LastName,
		FirstName:   input.FirstName,
		Patronymic:  input.Patronymic,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Password:    string(hashedPassword),
	}

	err = s.repo.CreateUser(c, newUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return newUser, nil
}

func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}

func GetUserByID(db *gorm.DB, id uint) (*models.User, error) {
	var user models.User
	err := db.First(&user, id).Error
	return &user, err
}

func GetUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	user := models.User{}
	err := db.Where("email = ?", email).First(&user).Error
	return &user, err
}
