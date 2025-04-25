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

func (service *UrlService) Create(userId uint, address string) (string, error) {
	url, err := service.UrlRepository.Create(&models.URL{
		UserID:   userId,
		Address:  address,
	})

	if err != nil {
		return "", err
	}

	return url.Address, nil
}

func (service *UrlService) Delete(id, userId uint) error {
	err := service.UrlRepository.Delete(id, userId)

	if err != nil {
		return err
	}

	return nil
}

func (service *UrlService) GetAll(userId uint) ([]models.URL, error) {
	urls, err := service.UrlRepository.FindByUser(userId)

	if err != nil {
		return nil, err
	}

	return urls, nil
}
