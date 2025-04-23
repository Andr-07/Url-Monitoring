package auth

import (
	"errors"
	"fmt"
	"go-monitoring/internal/models"
	"go-monitoring/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	exsistedUser, _ := service.UserRepository.FindByEmail(email)
	if exsistedUser != nil {
		return "", errors.New(ErrUserExists)
	}
	hashedPasowrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := &models.User{
		Email:    email,
		Password: string(hashedPasowrd),
		Name:     name,
	}
	_, err = service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}

func (service *AuthService) Login(email, password string) (string, error) {
	user, _ := service.UserRepository.FindByEmail(email)
	fmt.Println(user)
	if user == nil {
		return "", errors.New(ErrWrongCredentials)
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrWrongCredentials)
	}

	return user.Email, nil
}

