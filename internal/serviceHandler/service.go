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
	DeleteServiceByID(id uint) error
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

func (s Service) UpdateServiceByID(serviceID uint, serviceToUpdate entity.Service) (entity.Service, error) {
	service, err := s.repo.GetServiceByID(serviceID)
	if err != nil {
		return entity.Service{}, err
	}

	// Update the service fields
	if serviceToUpdate.Name != nil {
		service.Name = serviceToUpdate.Name
	}
	if serviceToUpdate.Description != "" {
		service.Description = serviceToUpdate.Description
	}

	updatedService, err := s.repo.UpdateService(service)
	if err != nil {
		return entity.Service{}, err
	}

	return updatedService, nil
}

func (s Service) DeleteServiceByID(id uint) error {
	err := s.repo.DeleteServiceByID(id)
	if err != nil {
		return err
	}

	return nil
}
