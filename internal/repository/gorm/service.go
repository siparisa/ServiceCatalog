package repository

import (
	"errors"
	"github.com/siparisa/ServiceCatalog/internal/controller/helper/request"
	"github.com/siparisa/ServiceCatalog/internal/entity"
	"gorm.io/gorm"
	"strings"
)

type IDataService interface {
	GetServices(servicesToGet entity.Service, pagination request.PaginationSettings) ([]entity.Service, error)
	GetServiceByID(id uint) (entity.Service, error)
	GetVersionsByServiceID(serviceID uint) ([]entity.Version, error)
}

type Service struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) IDataService {
	return &Service{
		db: db,
	}
}

func (r *Service) GetServices(servicesToGet entity.Service, pagination request.PaginationSettings) ([]entity.Service, error) {
	var services []entity.Service
	query := r.db.Table("services")

	if servicesToGet.Name != nil {
		// Use the ILIKE operator for case-insensitive partial match
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+strings.ToLower(*servicesToGet.Name)+"%")
	}

	if servicesToGet.Description != "" {
		// Use the ILIKE operator for case-insensitive partial match
		query = query.Or("LOWER(description) LIKE LOWER(?)", "%"+strings.ToLower(servicesToGet.Description)+"%")
	}

	// Apply pagination parameters
	offset := (pagination.Page - 1) * pagination.Limit
	query = query.Offset(offset).Limit(pagination.Limit)

	err := query.Find(&services).Error
	if err != nil {
		return nil, err
	}

	for i := range services {
		versions, err := r.GetVersionsByServiceID(services[i].ID)
		if err != nil {
			return nil, err
		}

		var entityVersions []entity.Version
		for _, v := range versions {
			entityVersions = append(entityVersions, entity.Version{
				Model:     v.Model,
				ServiceID: v.ServiceID,
				Version:   v.Version,
			})
		}

		services[i].Versions = entityVersions
	}

	return services, nil
}

func (r *Service) GetServiceByID(id uint) (entity.Service, error) {
	var service entity.Service
	err := r.db.Table("services").First(&service, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Service{}, gorm.ErrRecordNotFound
		}
		return entity.Service{}, err
	}
	return service, nil
}

func (r *Service) GetVersionsByServiceID(serviceID uint) ([]entity.Version, error) {
	var versions []entity.Version
	err := r.db.Table("versions").Where("service_id = ?", serviceID).Find(&versions).Error
	if err != nil {
		return nil, err
	}
	return versions, nil
}
