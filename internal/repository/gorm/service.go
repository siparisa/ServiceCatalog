package repository

import (
	"github.com/siparisa/ServiceCatalog/internal/entity"
	"gorm.io/gorm"
)

type IService interface {
	GetServices(servicesToGet entity.Service) ([]entity.Service, error)
	GetServiceByID(id uint) (entity.Service, error)
}

type Service struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) Service {
	return Service{
		db: db,
	}
}

func (r *Service) GetServices(servicesToGet entity.Service) ([]entity.Service, error) {
	var services []entity.Service
	err := r.db.Find(&services).Error
	if err != nil {
		return nil, err
	}
	return services, nil
}

func (r *Service) GetServiceByID(id uint) (entity.Service, error) {
	var service entity.Service
	err := r.db.First(&service, id).Error
	if err != nil {
		return entity.Service{}, err
	}
	return service, nil
}
