package request

// GetServicesQueryParams is the request query params for getting a list of services
type GetServicesQueryParams struct {
	Name        *string `form:"name" example:"Service Name"`
	Description *string `form:"description" example:"Description Of a Service"`
	Page        int     `form:"page" example:"1"`
	Limit       int     `form:"limit" example:"10"`
}

// ServiceURI contains the uri parts for service requests
type ServiceURI struct {
	ServiceID uint `uri:"serviceID" binding:"required"`
}

// CreateServiceBody is the request body for creating a service
type CreateServiceBody struct {
	Data CreateService `json:"data" binding:"required"`
}

// CreateService is the request data for creating a service
type CreateService struct {
	Name        string `json:"name" binding:"required" example:"service1"`
	Description string `json:"description" binding:"required" example:"My Service"`
	Version     string `json:"version" binding:"required" example:"1.0"`
}

// UpdateServiceBody is the request body for updating a service
type UpdateServiceBody struct {
	Data UpdateService `json:"data" binding:"required"`
}

// UpdateService is the request data for updating a service
type UpdateService struct {
	Name        string `json:"name" binding:"omitempty,min=1" example:"service1"`
	Description string `json:"description" binding:"omitempty,min=1" example:"My Service"`
}
