// Inside serviceHandler/serviceHandler.go

package serviceHandler

import (
	"github.com/siparisa/ServiceCatalog/internal/entity"
	"github.com/siparisa/ServiceCatalog/internal/repository"
)

// GetServices retrieves a paginated list of services.
func GetServices(servicesToGet entity.Service, page, limit string) (entity.Service, error) {
	// Call the repository layer to retrieve a paginated list of services
	services, err := repository.GetServices
	if err != nil {
		return nil, err
	}

	return services, nil
}

// GetServiceByID retrieves a serviceHandler by its ID.
func GetServiceByID(id string) (*model.Service, error) {
	// Call the repository layer to retrieve the serviceHandler by ID
	service, err := repository.GetServiceByID(id)
	if err != nil {
		return nil, err
	}

	return service, nil
}
