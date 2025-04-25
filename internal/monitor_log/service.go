package monitor_log

import (
	"fmt"
	"go-monitoring/internal/models"
	"go-monitoring/internal/repository"
	"go-monitoring/pkg/notifier"
	"log"
	"net/http"
	"time"
)

type MonitorLogService struct {
	MonitorLogRepository *repository.MonitorLogRepository
	UrlRepository        *repository.UrlRepository
	Notifier             notifier.Notifier
	stopChan             chan struct{}
}

func NewMonitorLogService(
	monitorLogRepository *repository.MonitorLogRepository,
	urlRepository *repository.UrlRepository,
	Notifier notifier.Notifier,
	stopChan chan struct{},
) *MonitorLogService {
	return &MonitorLogService{
		MonitorLogRepository: monitorLogRepository,
		UrlRepository:        urlRepository,
		Notifier:             Notifier,
		stopChan:             stopChan,
	}
}

func (service *MonitorLogService) Start() {
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				service.checkUrls()

			case <-service.stopChan:
				log.Println("Stopping monitor service")
				return
			}

		}
	}()
	go service.cleanupOldLogs()
}

func (service *MonitorLogService) checkUrls() {
	urls, err := service.UrlRepository.GetAll()
	if err != nil {
		log.Println("error getting urls:", err)
		return
	}

	for _, u := range urls {
		go service.checkOne(u)
	}
}

func (service *MonitorLogService) checkOne(url models.URL) {
	start := time.Now()
	resp, err := http.Get(url.Address)
	status := models.StatusOK
	code := 200
	msg := ""

	if err != nil {
		status = models.StatusFail
		msg = err.Error()
		code = 0
	} else {
		code = resp.StatusCode
		if resp.StatusCode >= 400 {
			status = models.StatusFail
			msg = resp.Status
		}
		resp.Body.Close()
	}

	if status == models.StatusFail {
		err := service.Notifier.SendAlert(
			fmt.Sprintf("🚨 URL %s недоступен! Код: %d. Ошибка: %s", url.Address, code, msg),
		)
		if err != nil {
			log.Printf("Не удалось отправить Telegram-уведомление: %v", err)
		}
	}

	service.Create(MonitorLogDto{
		URLID:     url.ID,
		Timestamp: start,
		Status:    status,
		HTTPCode:  code,
		Error:     msg,
	})
}

func (service *MonitorLogService) Create(dto MonitorLogDto) (*models.MonitorLog, error) {
	url, err := service.MonitorLogRepository.Save(&models.MonitorLog{
		URLID:     dto.URLID,
		Timestamp: dto.Timestamp,
		Status:    dto.Status,
		HTTPCode:  dto.HTTPCode,
		Error:     dto.Error,
	})

	if err != nil {
		return nil, err
	}

	return url, nil
}

func (service *MonitorLogService) GetAll(urlId uint) ([]models.MonitorLog, error) {
	monitorLogs, err := service.MonitorLogRepository.FindByUrl(urlId)

	if err != nil {
		return nil, err
	}

	return monitorLogs, nil
}

func (service *MonitorLogService) cleanupOldLogs() {
	for {
		select {
		case <-time.After(1 * time.Hour):
			if err := service.MonitorLogRepository.DeleteOldLogs(); err != nil {
				log.Println("failed to delete old logs:", err)
			}
		case <-service.stopChan:
			return
		}
	}
}
