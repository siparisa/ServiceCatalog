package serviceHandler

import (
	"github.com/siparisa/ServiceCatalog/internal/controller/helper/request"
	"github.com/siparisa/ServiceCatalog/internal/entity"
	repository "github.com/siparisa/ServiceCatalog/internal/repository/gorm"
)

type IService interface {
	CreateService(service entity.Service) (entity.Service, error)
	GetServices(servicesToGet entity.Service, pagination request.PaginationSettings) ([]entity.Service, error)
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

func (s Service) CreateService(service entity.Service) (entity.Service, error) {
	createdService, err := s.repo.CreateService(service)
	if err != nil {
		return entity.Service{}, err
	}

	return createdService, nil
}

func (s Service) GetServices(servicesToGet entity.Service, pagination request.PaginationSettings) ([]entity.Service, error) {

	services, err := s.repo.GetServices(servicesToGet, pagination)
	if err != nil {
		return nil, err
	}

	return services, nil
}

func (s Service) GetServiceByID(id uint) (entity.Service, error) {
	return s.repo.GetServiceByID(id)
}
