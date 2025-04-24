package repository

import (
	"go-monitoring/internal/models"
	"go-monitoring/pkg/db"
)

type MonitorLogRepository struct {
	Database *db.Db
}

func NewMonitorLogRepository(database *db.Db) *MonitorLogRepository {
	return &MonitorLogRepository{
		Database: database,
	}
}

func (repo *MonitorLogRepository) Save(log *models.MonitorLog) (*models.MonitorLog, error) {
	result := repo.Database.DB.Create(log)
	if result.Error != nil {
		return nil, result.Error
	}
	return log, nil
}

func (repo *MonitorLogRepository) FindByUrl(urlId uint) ([]models.MonitorLog, error) {
	var logs []models.MonitorLog
	result := repo.Database.DB.Where("url_id = ?", urlId).Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}
	return logs, nil
}
