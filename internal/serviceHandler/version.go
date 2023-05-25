package serviceHandler

import (
	"github.com/siparisa/ServiceCatalog/internal/entity"
	repository "github.com/siparisa/ServiceCatalog/internal/repository/gorm"
)

type IVersion interface {
	CreateVersion(version entity.Version) (entity.Version, error)
	GetVersionsByServiceID(serviceID uint) ([]entity.Version, error)
}

type VersionService struct {
	repo repository.IVersionRepository
}

func NewVersion(repo repository.IVersionRepository) IVersion {
	return &VersionService{
		repo: repo,
	}
}

// CreateVersion creates a version
func (vs *VersionService) CreateVersion(version entity.Version) (entity.Version, error) {
	createdVersion, err := vs.repo.CreateVersion(version)
	if err != nil {
		return entity.Version{}, err
	}

	return createdVersion, nil
}

// GetVersionsByServiceID gets versions for a service
func (vs *VersionService) GetVersionsByServiceID(serviceID uint) ([]entity.Version, error) {
	return vs.repo.GetVersionsByServiceID(serviceID)
}
