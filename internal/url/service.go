package url

import (
	"go-monitoring/internal/models"
	"go-monitoring/internal/repository"
)

type UrlService struct {
	UrlRepository *repository.UrlRepository
}

func NewUrlService(urlRepository *repository.UrlRepository) *UrlService {
	return &UrlService{UrlRepository: urlRepository}
}

func (service *UrlService) Create(userId uint, address string, interval int) (string, error) {
	url, err := service.UrlRepository.Create(&models.URL{
		UserID: userId,
		Address: address,
		Interval: interval,
	})

	if err != nil {
		return "", err
	}

	return url.Address, nil
}

