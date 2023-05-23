package serviceHandler

import (
	"github.com/siparisa/ServiceCatalog/internal/entity"
	repository "github.com/siparisa/ServiceCatalog/internal/repository/gorm"
)

type IService interface {
	GetServices(servicesToGet entity.Service, page, limit int) ([]entity.Service, error)
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

func (s Service) GetServices(servicesToGet entity.Service, page, limit int) ([]entity.Service, error) {

	services, err := s.repo.GetServices(servicesToGet, page, limit)
	if err != nil {
		return nil, err
	}

	//// Fetch versions for each service
	//for i := range services {
	//	versions, err := s.repo.GetVersionsByServiceID(services[i].ID)
	//	if err != nil {
	//		return nil, err
	//	}
	//	services[i].Versions = versions
	//}

	return services, nil
}

func (s Service) GetServiceByID(id uint) (entity.Service, error) {
	return s.repo.GetServiceByID(id)
}
