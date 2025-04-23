package main

import (
	"go-monitoring/config"
	"go-monitoring/internal/auth"
	"go-monitoring/internal/repository"
	"go-monitoring/pkg/db"
	"log"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	userRepository := repository.NewUserRepository(db)
	authService := auth.NewAuthService(userRepository)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}