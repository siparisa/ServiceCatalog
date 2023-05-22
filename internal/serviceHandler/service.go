package serviceHandler

import (
	"github.com/siparisa/ServiceCatalog/internal/entity"
	repository "github.com/siparisa/ServiceCatalog/internal/repository/gorm"
)

type IService interface {
	GetServices(servicesToGet entity.Service) ([]entity.Service, error)
	GetServiceByID(id uint) (entity.Service, error)
}

type Service struct {
	repo repository.IDataService
}

func NewService(repo repository.IDataService) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) GetServices(servicesToGet entity.Service) ([]entity.Service, error) {
	return s.repo.GetServices(servicesToGet)
}

func (s Service) GetServiceByID(id uint) (entity.Service, error) {
	return s.repo.GetServiceByID(id)
}
