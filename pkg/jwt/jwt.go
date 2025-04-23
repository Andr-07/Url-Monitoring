package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct {
	UserID uint
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(data JWTData) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": data.UserID,
	})
	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *JWT) Parse(token string) (bool, *JWTData) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	claims, _ := t.Claims.(jwt.MapClaims)
	if v, ok := claims["userId"].(float64); ok {
		return true, &JWTData{UserID: uint(v)}
	}
	return false, nil
}
