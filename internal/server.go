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
		services.GET("", func(c *gin.Context) {
			controller.GetServices(db, c)
		})
		services.GET("/:id", func(c *gin.Context) {
			controller.GetServiceByID(db, c)
		})
	}

	return r
}
