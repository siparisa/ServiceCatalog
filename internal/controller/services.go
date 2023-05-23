package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/siparisa/ServiceCatalog/internal/controller/helper/request"
	"github.com/siparisa/ServiceCatalog/internal/controller/helper/response"
	"github.com/siparisa/ServiceCatalog/internal/entity"
	repository "github.com/siparisa/ServiceCatalog/internal/repository/gorm"
	"github.com/siparisa/ServiceCatalog/internal/serviceHandler"
	"gorm.io/gorm"
	"strconv"
)

func GetServices(db *gorm.DB, c *gin.Context) {
	var qp request.GetServicesQueryParams
	if err := c.ShouldBindQuery(&qp); err != nil {
		response.BadRequest(c, "Invalid query parameters", err.Error())
		return
	}

	pagination := request.PaginationSettings{
		Page:  qp.Page,
		Limit: qp.Limit,
	}

	// Check if Page is null or 0, assign default value
	if pagination.Page == 0 {
		pagination.Page = 1
	}

	// Check if Limit is null or 0, assign default value
	if pagination.Limit == 0 {
		pagination.Limit = 10
	}

	servicesToGet := entity.Service{
		Name:        qp.Name,
		Description: qp.Description,
	}

	// Create a repository instance
	repo := repository.NewServiceRepository(db)

	// Create a service instance
	serviceHndlr := serviceHandler.NewService(repo)

	// Call the service layer to retrieve a paginated list of services
	services, err := serviceHndlr.GetServices(servicesToGet, pagination)
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve services", err.Error())
		return
	}

	response.OK(c, services)
}

func GetServiceByID(db *gorm.DB, c *gin.Context) {
	var uri request.ServiceURI
	if err := c.ShouldBindUri(&uri); err != nil {
		response.BadRequest(c, "Missing ID", err.Error())
		return
	}
	serviceID, err := strconv.ParseUint(uri.ServiceID, 10, 64)
	if err != nil {
		response.BadRequest(c, "Invalid service ID", err.Error())
		return
	}

	// Create a repository instance
	repo := repository.NewServiceRepository(db)

	// Create a service instance
	serviceHndlr := serviceHandler.NewService(repo)

	// Call the service layer to retrieve the service by ID
	service, err := serviceHndlr.GetServiceByID(uint(serviceID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NotFound(c, "Service not found")
		} else {
			response.InternalServerError(c, "Failed to retrieve service", err.Error())
		}
		return
	}

	response.OK(c, service)
}

func CreateService(db *gorm.DB, c *gin.Context) {

	var body request.CreateServiceBody
	if err := c.ShouldBindJSON(&body); err != nil {
		response.BuildErrorResponse(c, "Invalid request payload", err.Error())
		return
	}

	serviceToCreate := entity.Service{
		Name:        &body.Data.Name,
		Description: body.Data.Description,
	}

	// Create a repository instance
	repo := repository.NewServiceRepository(db)

	// Create a service instance
	serviceHndlr := serviceHandler.NewService(repo)

	// Call the service layer to create a new service
	createdService, err := serviceHndlr.CreateService(serviceToCreate)
	if err != nil {
		response.InternalServerError(c, "Failed to create service", err.Error())
		return
	}

	response.OK(c, createdService)
}
