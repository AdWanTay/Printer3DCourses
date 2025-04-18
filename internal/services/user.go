package services

import (
	"Printer3DCourses/internal/dto"
	"Printer3DCourses/internal/models"
	"Printer3DCourses/internal/repositories"
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
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

	validate := validator.New()
	err = validate.Struct(newUser)
	if err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		errorString := "failed to validate user fields: \n"
		for _, validationError := range validationErrors {
			errorString += validationError.Field() + "\n"
		}
		return nil, fmt.Errorf(errorString)
	}

	err = s.repo.CreateUser(c, newUser)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return newUser, nil
}

func (s *UserService) LoginUser(c context.Context, input dto.LoginRequest) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(c, input.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return user, nil
}
