package request

// GetServicesQueryParams is the request query params for getting a list of services
type GetServicesQueryParams struct {
	Name *string `form:"name" example:"Service Name"`
} // @name Request.GetServicesQueryParams
