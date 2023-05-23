package controller

// Inside controller/services.go

import (
	"github.com/gin-gonic/gin"
	"github.com/siparisa/ServiceCatalog/internal/controller/helper/request"
	"github.com/siparisa/ServiceCatalog/internal/entity"
	repository "github.com/siparisa/ServiceCatalog/internal/repository/gorm"
	"github.com/siparisa/ServiceCatalog/internal/serviceHandler"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetServices(db *gorm.DB, c *gin.Context) {
	var qp request.GetServicesQueryParams
	if err := c.ShouldBindQuery(&qp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	// Extract pagination parameters from query string
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1 // Set a default page number if not provided or invalid
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10 // Set a default limit if not provided or invalid
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
	services, err := serviceHndlr.GetServices(servicesToGet, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve services"})
		return
	}

	c.JSON(http.StatusOK, services)
}

func GetServiceByID(db *gorm.DB, c *gin.Context) {

	var uri request.ServiceURI
	if err := c.ShouldBindUri(&uri); err != nil {
		// tatus, res := response.BuildErrorResponse(ctx, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
		return
	}
	serviceID, err := strconv.ParseUint(uri.ServiceID, 10, 64)

	// Create a repository instance
	repo := repository.NewServiceRepository(db)

	// Create a service instance
	serviceHndlr := serviceHandler.NewService(repo)

	// Call the service layer to retrieve the service by ID
	service, err := serviceHndlr.GetServiceByID(uint(serviceID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	c.JSON(http.StatusOK, service)

}
