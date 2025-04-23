package repository

import (
	"go-monitoring/pkg/db"
	"go-monitoring/internal/models"
)


type UrlRepository struct {
	Database *db.Db
}

func NewUrlRepository(database *db.Db) *UrlRepository {
	return &UrlRepository{
		Database: database,
	}
}

func (repo *UrlRepository) Create(url *models.URL) (*models.URL, error) {
	result := repo.Database.DB.Create(url)
	if result.Error != nil {
		return nil, result.Error
	}
	return url, nil
}

func (repo *UrlRepository) Delete(url *models.URL) (*models.URL, error) {
	result := repo.Database.DB.Delete(url)
	if result.Error != nil {
		return nil, result.Error
	}
	return url, nil
}


