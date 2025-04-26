package mock

import (
	"errors"
	"go-monitoring/internal/auth"
	"go-monitoring/internal/models"
)

type MockUserRepository struct {
	users map[string]*models.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[string]*models.User),
	}
}

func (repo *MockUserRepository) FindByEmail(email string) (*models.User, error) {
	user, exists := repo.users[email]
	if exists {
		return user, nil
	}
	return nil, nil
}

func (repo *MockUserRepository) Create(user *models.User) (*models.User, error) {
	if _, exists := repo.users[user.Email]; exists {
		return nil, errors.New(auth.ErrUserExists)
	}
	user.ID = uint(len(repo.users) + 1)
	repo.users[user.Email] = user
	return user, nil
}
