package controller

// Inside controller/services.go

import (
	"github.com/gin-gonic/gin"
	"github.com/siparisa/ServiceCatalog/internal/controller/helper/request"
	"github.com/siparisa/ServiceCatalog/internal/entity"
	"github.com/siparisa/ServiceCatalog/internal/serviceHandler"
	"net/http"
)

func GetServices(c *gin.Context) {
	// Extract pagination parameters from query string
	page := c.Query("page")
	limit := c.Query("limit")

	var qp request.GetServicesQueryParams
	if err := c.ShouldBindQuery(&qp); err != nil {
		status, res := response.BuildErrorResponse(ctx, err)
		c.JSON(status, res)
		return
	}
	servicesToGet := entity.Service{
		Name: qp.Name,
	}

	// pgnSettings := request.GetPaginationSettings(ctx)

	// Call the serviceHandler layer to retrieve a paginated list of services
	services, err := serviceHandler.GetServices(servicesToGet, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve services"})
		return
	}

	c.JSON(http.StatusOK, services)
}

func GetServiceByID(c *gin.Context) {
	// Extract the serviceHandler ID from the URL parameter
	serviceID := c.Param("id")

	// Call the serviceHandler layer to retrieve the serviceHandler by ID
	service, err := serviceHandler.GetServiceByID(serviceID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	c.JSON(http.StatusOK, service)
}
