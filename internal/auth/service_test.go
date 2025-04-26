package auth_test

import (
	"go-monitoring/internal/auth"
	"go-monitoring/internal/mock/repository"
	"go-monitoring/internal/models"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestRegisterSuccess(t *testing.T) {
	const initialEmail = "test@test.ru"
	mockRepo := mock.NewMockUserRepository()
	authService := auth.NewAuthService(mockRepo)

	userId, err := authService.Register(initialEmail, "1", "Test")
	if err != nil {
		t.Fatal(err)
	}
	if userId != 1 {
		t.Fatalf("UserId %d do not match 1", userId)
	}
}

func TestRegisterUserExists(t *testing.T) {
	const initialEmail = "test@test.ru"
	mockRepo := mock.NewMockUserRepository()
	authService := auth.NewAuthService(mockRepo)

	_, err := authService.Register(initialEmail, "1", "Test")
	if err != nil {
		t.Fatal(err)
	}

	_, err = authService.Register(initialEmail, "1", "Test")
	if err == nil {
		t.Fatal("Expected error when registering existing user")
	}
	if err.Error() != auth.ErrUserExists {
		t.Fatalf("Expected 'user already exists' error, got %v", err)
	}
}

func TestLoginSuccess(t *testing.T) {
	const email = "test@test.ru"
	const password = "password123"

	mockRepo := mock.NewMockUserRepository()
	authService := auth.NewAuthService(mockRepo)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}

	mockRepo.Create(&models.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     "Test User",
	})

	userId, err := authService.Login(email, password)
	if err != nil {
		t.Fatal(err)
	}
	if userId == 0 {
		t.Fatal("Expected valid user ID, got 0")
	}
}

func TestLoginWrongPassword(t *testing.T) {
	const email = "test@test.ru"
	const correctPassword = "password123"
	const wrongPassword = "wrongpassword"

	mockRepo := mock.NewMockUserRepository()
	authService := auth.NewAuthService(mockRepo)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(correctPassword), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}

	mockRepo.Create(&models.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     "Test User",
	})

	_, err = authService.Login(email, wrongPassword)
	if err == nil {
		t.Fatal("Expected error for wrong password")
	}
	if err.Error() != "wrong credentials" {
		t.Fatalf("Expected 'wrong credentials' error, got %v", err)
	}
}

func TestLoginUserNotFound(t *testing.T) {
	const email = "nonexistent@test.ru"
	const password = "password123"

	mockRepo := mock.NewMockUserRepository()
	authService := auth.NewAuthService(mockRepo)

	_, err := authService.Login(email, password)
	if err == nil {
		t.Fatal("Expected error for non-existing user")
	}
	if err.Error() != "wrong credentials" {
		t.Fatalf("Expected 'wrong credentials' error, got %v", err)
	}
}