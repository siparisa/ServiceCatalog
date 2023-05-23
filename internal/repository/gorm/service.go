package repository

import (
	"github.com/siparisa/ServiceCatalog/internal/entity"
	"gorm.io/gorm"
	"strings"
)

type IDataService interface {
	GetServices(servicesToGet entity.Service, page, limit int) ([]entity.Service, error)
	GetServiceByID(id uint) (entity.Service, error)
}

type Service struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) IDataService {
	return &Service{
		db: db,
	}
}

func (r *Service) GetServices(servicesToGet entity.Service, page, limit int) ([]entity.Service, error) {
	var services []entity.Service
	query := r.db.Table("services")

	if servicesToGet.Name != nil {
		// Use the ILIKE operator for case-insensitive partial match
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+strings.ToLower(*servicesToGet.Name)+"%")
	}

	// Apply pagination parameters
	offset := (page - 1) * limit
	query = query.Offset(offset).Limit(limit)

	err := query.Find(&services).Error
	if err != nil {
		return nil, err
	}
	return services, nil
}

func (r *Service) GetServiceByID(id uint) (entity.Service, error) {
	var service entity.Service
	err := r.db.Table("services").First(&service, id).Error
	if err != nil {
		return entity.Service{}, err
	}
	return service, nil
}
