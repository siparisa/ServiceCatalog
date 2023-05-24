package repository

import (
	"github.com/siparisa/ServiceCatalog/internal/entity"
	"gorm.io/gorm"
)

type IVersionRepository interface {
	CreateVersion(version entity.Version) (entity.Version, error)
	GetVersionsByServiceID(serviceID uint) ([]entity.Version, error)
}

type VersionRepository struct {
	db *gorm.DB
}

func NewVersionRepository(db *gorm.DB) IVersionRepository {
	return &VersionRepository{
		db: db,
	}
}

func (r *VersionRepository) CreateVersion(version entity.Version) (entity.Version, error) {
	err := r.db.Table("versions").Create(&version).Error
	if err != nil {
		return entity.Version{}, err
	}

	return version, nil
}

func (r *VersionRepository) GetVersionsByServiceID(serviceID uint) ([]entity.Version, error) {
	var versions []entity.Version
	err := r.db.Table("versions").Where("service_id = ?", serviceID).Find(&versions).Error
	if err != nil {
		return nil, err
	}
	return versions, nil
}
