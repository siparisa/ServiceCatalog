package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/siparisa/ServiceCatalog/internal/controller"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	services := r.Group("/services")
	{
		services.POST("", func(c *gin.Context) { controller.CreateService(db, c) })
		services.GET("", func(c *gin.Context) { controller.GetServices(db, c) })
		services.GET("/:serviceID", func(c *gin.Context) { controller.GetServiceByID(db, c) })
		services.PUT("/:serviceID", func(c *gin.Context) { controller.UpdateServiceByID(db, c) })
		services.DELETE("/:serviceID", func(c *gin.Context) { controller.DeleteServiceByID(db, c) })

		// to implement in future
		// services.GET("/:id/versions", func(c *gin.Context) { controller.GetServiceVersions(db, c) })
		// services.POST("/:id/versions", func(c *gin.Context) { controller.CreateServiceVersion(db, c) })

	}

	return r
}
