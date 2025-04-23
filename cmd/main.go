package main

import (
	"go-monitoring/config"
	"go-monitoring/internal/auth"
	"go-monitoring/internal/repository"
	"go-monitoring/internal/url"
	"go-monitoring/pkg/db"
	"log"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	userRepository := repository.NewUserRepository(db)
	urlRepository := repository.NewUrlRepository(db)

	// Services
	authService := auth.NewAuthService(userRepository)
	urlService := url.NewUrlService(urlRepository)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	url.NewUrlHandler(router, url.UrlHandlerDeps{
		Config:     conf,
		UrlService: urlService,
	})

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
