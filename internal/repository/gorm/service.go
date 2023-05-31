package repository

import (
	"errors"
	"github.com/siparisa/ServiceCatalog/internal/controller/helper/request"
	"github.com/siparisa/ServiceCatalog/internal/entity"
	"gorm.io/gorm"
	"strings"
)

// IServiceRepository is an interface for service repository
type IServiceRepository interface {
	CreateService(service entity.Service) (entity.Service, error)
	GetServices(servicesToGet entity.Service, pagination request.PaginationSettings) ([]entity.Service, error)
	GetServiceByID(id uint) (entity.Service, error)
	GetVersionsByServiceID(serviceID uint) ([]entity.Version, error)
	UpdateService(service entity.Service) (entity.Service, error)
	DeleteServiceByID(id uint) error
}

type Service struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) IServiceRepository {
	return &Service{
		db: db,
	}
}

// CreateService creates a service
func (r *Service) CreateService(service entity.Service) (entity.Service, error) {
	err := r.db.Table("services").Create(&service).Error
	if err != nil {
		return entity.Service{}, err
	}

	return service, nil
}

// GetServices gets all services or based on the given parameters
func (r *Service) GetServices(servicesToGet entity.Service, pagination request.PaginationSettings) ([]entity.Service, error) {
	var services []entity.Service
	query := r.db.Table("services").Order("created_at DESC")

	if servicesToGet.Name != nil {
		// Use the LIKE operator for case-insensitive partial match
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+strings.ToLower(*servicesToGet.Name)+"%")
	}

	if servicesToGet.Description != nil {
		// Use the LIKE operator for case-insensitive partial match
		query = query.Or("LOWER(description) LIKE LOWER(?)", "%"+strings.ToLower(*servicesToGet.Description)+"%")
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

// GetServiceByID gets service by ID
func (r *Service) GetServiceByID(id uint) (entity.Service, error) {
	var service entity.Service
	err := r.db.Table("services").First(&service, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Service{}, gorm.ErrRecordNotFound
		}
		return entity.Service{}, err
	}

	versions, err := r.GetVersionsByServiceID(service.ID)
	if err != nil {
		return entity.Service{}, err
	}

	service.Versions = versions
	return service, nil
}

// GetVersionsByServiceID gets versions by serviceID
func (r *Service) GetVersionsByServiceID(serviceID uint) ([]entity.Version, error) {
	var versions []entity.Version
	err := r.db.Table("versions").Where("service_id = ?", serviceID).Find(&versions).Error
	if err != nil {
		return nil, err
	}
	return versions, nil
}

// UpdateService updates a service
func (r *Service) UpdateService(service entity.Service) (entity.Service, error) {
	err := r.db.Table("services").Save(&service).Error
	if err != nil {
		return entity.Service{}, err
	}

	return service, nil
}

// DeleteServiceByID deletes a service by ID
func (r *Service) DeleteServiceByID(id uint) error {
	// Check if the service exists
	existingService, err := r.GetServiceByID(id)
	if err != nil {
		return err
	}

	// Delete the associated version records
	err = r.db.Table("versions").Where("service_id = ?", id).Delete(&entity.Version{}).Error
	if err != nil {
		return err
	}

	// Delete the service record
	err = r.db.Table("services").Delete(&existingService).Error
	if err != nil {
		return err
	}

	return nil
}
