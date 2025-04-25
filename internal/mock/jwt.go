package mock

import (
	"go-monitoring/pkg/jwt"
)

type MockJWTService struct{}

func (m *MockJWTService) Create(data jwt.JWTData) (string, error) {
	return "mocked_token", nil
}

func NewMockJWTService() *jwt.JWT {
	return &jwt.JWT{}
}
