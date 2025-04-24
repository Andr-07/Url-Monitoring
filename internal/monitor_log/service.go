package monitor_log

import (
	"go-monitoring/internal/models"
	"go-monitoring/internal/repository"
)

type MonitorLogService struct {
	MonitorLogRepository *repository.MonitorLogRepository
}

func NewMonitorLogService(monitorLogRepository *repository.MonitorLogRepository) *MonitorLogService {
	return &MonitorLogService{MonitorLogRepository: monitorLogRepository}
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
