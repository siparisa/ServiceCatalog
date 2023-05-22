package controller

// Inside controller/services.go

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/siparisa/ServiceCatalog/internal/controller/helper/request"
	"github.com/siparisa/ServiceCatalog/internal/entity"
	repository "github.com/siparisa/ServiceCatalog/internal/repository/gorm"
	"github.com/siparisa/ServiceCatalog/internal/serviceHandler"
	"gorm.io/gorm"
	"net/http"
)

func GetServices(db *gorm.DB, c *gin.Context) {
	// Extract pagination parameters from query string
	// page := c.Query("page")
	// limit := c.Query("limit")

	var qp request.GetServicesQueryParams
	if err := c.ShouldBindQuery(&qp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	servicesToGet := entity.Service{
		Name: qp.Name,
	}

	// Create a repository instance
	repo := repository.NewServiceRepository(db) // Replace 'db' with your Gorm DB instance

	// Create a service instance
	service := serviceHandler.NewService(repo)

	// Call the service layer to retrieve a paginated list of services
	services, err := service.GetServices(servicesToGet)
	if err != nil {
		fmt.Println("err::", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve services"})
		return
	}

	c.JSON(http.StatusOK, services)
}

func GetServiceByID(db *gorm.DB, c *gin.Context) {
	//// Extract the service ID from the URL parameter
	//serviceID := c.Param("id")
	//
	//// Create a repository instance
	//repo := repository.NewServiceRepository(db) // Replace 'db' with your Gorm DB instance
	//
	//// Create a service instance
	//service := serviceHandler.NewService(repo)
	//
	//// Call the service layer to retrieve the service by ID
	//service, err := service.GetServiceByID(serviceID)
	//if err != nil {
	//	c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
	//	return
	//}
	//
	//c.JSON(http.StatusOK, service)
	c.JSON(http.StatusOK, "gjhgjhghj")

}
