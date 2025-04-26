package auth_test

import (
	"bytes"
	"encoding/json"
	"go-monitoring/config"
	"go-monitoring/internal/auth"
	"go-monitoring/internal/repository"
	"go-monitoring/pkg/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootstrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		return nil, nil, err
	}

	userRepo := repository.NewUserRepository(&db.Db{
		DB: gormDb,
	})

	handler := auth.AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: auth.NewAuthService(userRepo),
	}
	return &handler, mock, nil
}

func TestHandlerLoginSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatal(err)
		return
	}

	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("test@test.ru", "$2a$10$rcu89fWNDrpFn/wgIbl7lunRLKOwkazymxvYa6MHFUeTSn2clbeuy$2a$10$rcu89fWNDrpFn/wgIbl7lunRLKOwkazymxvYa6MHFUeTSn2clbeuy")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test@test.ru",
		Password: "test",
	})

	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/auth/login", reader)

	handler.Login()(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("got %d, expected %d", w.Code, http.StatusOK)
	}
}

func TestHandlerLoginFailWrongPassword(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatal(err)
		return
	}

	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("test@test.ru", "$2a$10$rcu89fWNDrpFn/wgIbl7lunRLKOwkazymxvYa6MHFUeTSn2clbeuy")
	mock.ExpectQuery("SELECT email, password FROM users WHERE email = ?").WithArgs("test@test.ru").WillReturnRows(rows)

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test@test.ru",
		Password: "wrongpassword",
	})

	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/auth/login", reader)

	handler.Login()(w, r)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("got %d, expected %d", w.Code, http.StatusUnauthorized)
	}
}

func TestHandlerLoginFailUserNotFound(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatal(err)
		return
	}

	mock.ExpectQuery("SELECT email, password FROM users WHERE email = ?").WithArgs("notfound@test.ru").WillReturnRows(sqlmock.NewRows([]string{"email", "password"}))

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "notfound@test.ru",
		Password: "testpassword",
	})

	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/auth/login", reader)

	handler.Login()(w, r)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("got %d, expected %d", w.Code, http.StatusUnauthorized)
	}
}
