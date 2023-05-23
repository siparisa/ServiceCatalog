package request

// GetServicesQueryParams is the request query params for getting a list of services
type GetServicesQueryParams struct {
	Name        *string `form:"name" example:"Service Name"`
	Description string  `form:"description" example:"Description Of a Service"`
}

// ServiceURI contains the uri parts for service requests
type ServiceURI struct {
	ServiceID string `uri:"id" binding:"required"`
}
