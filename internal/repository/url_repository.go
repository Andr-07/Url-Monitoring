package repository

import (
	"go-monitoring/internal/models"
	"go-monitoring/pkg/db"
)

type UrlRepository struct {
	Database *db.Db
}

func NewUrlRepository(database *db.Db) *UrlRepository {
	return &UrlRepository{
		Database: database,
	}
}

func (repo *UrlRepository) FindByUser(userId uint) ([]models.URL, error) {
	var urls []models.URL
	result := repo.Database.DB.Where("user_id = ?", userId).Find(&urls)
	if result.Error != nil {
		return nil, result.Error
	}
	return urls, nil
}

func (repo *UrlRepository) Create(url *models.URL) (*models.URL, error) {
	result := repo.Database.DB.Create(url)
	if result.Error != nil {
		return nil, result.Error
	}
	return url, nil
}

func (repo *UrlRepository) Delete(id, userId uint) error {
	result := repo.Database.DB.
		Where("id = ? AND user_id = ?", id, userId).
		Delete(&models.URL{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
