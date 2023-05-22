package request

// GetServicesQueryParams is the request query params for getting a list of services
type GetServicesQueryParams struct {
	Name *string `form:"name" example:"Service Name"`
}

// ServiceURI contains the uri parts for service requests
type ServiceURI struct {
	ServiceID string `uri:"id" binding:"required"`
}
