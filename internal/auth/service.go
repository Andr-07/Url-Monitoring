package auth

import (
	"errors"
	"go-monitoring/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}

type AuthService struct {
	UserRepository UserRepository
}

func NewAuthService(userRepository UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Register(email, password, name string) (uint, error) {
	exsistedUser, _ := service.UserRepository.FindByEmail(email)
	if exsistedUser != nil {
		return 0, errors.New(ErrUserExists)
	}
	hashedPasowrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user := &models.User{
		Email:    email,
		Password: string(hashedPasowrd),
		Name:     name,
	}
	_, err = service.UserRepository.Create(user)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (service *AuthService) Login(email, password string) (uint, error) {
	user, _ := service.UserRepository.FindByEmail(email)
	if user == nil {
		return 0, errors.New(ErrWrongCredentials)
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return 0, errors.New(ErrWrongCredentials)
	}

	return user.ID, nil
}

