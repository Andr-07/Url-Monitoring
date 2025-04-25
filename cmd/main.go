package main

import (
	"go-monitoring/config"
	"go-monitoring/internal/auth"
	"go-monitoring/internal/monitor_log"
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
	stopChan := make(chan struct{})

	// Repositories
	userRepository := repository.NewUserRepository(db)
	urlRepository := repository.NewUrlRepository(db)
	monitorLogRepository := repository.NewMonitorLogRepository(db)

	// Services
	authService := auth.NewAuthService(userRepository)
	urlService := url.NewUrlService(urlRepository)
	monitorLogService := monitor_log.NewMonitorLogService(monitorLogRepository, urlRepository, stopChan)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	url.NewUrlHandler(router, url.UrlHandlerDeps{
		Config:     conf,
		UrlService: urlService,
	})
	monitor_log.NewMonitorLogHandler(router, monitor_log.MonitorLogHandlerDeps{
		Config:            conf,
		MonitorLogService: monitorLogService,
	})

	log.Println("Server started on :8080")
	monitorLogService.Start()
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
