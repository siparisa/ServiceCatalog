package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/siparisa/ServiceCatalog/internal/controller/helper/request"
	"github.com/siparisa/ServiceCatalog/internal/controller/helper/response"
	repository "github.com/siparisa/ServiceCatalog/internal/repository/gorm"
	"github.com/siparisa/ServiceCatalog/internal/serviceHandler"
	"gorm.io/gorm"
)

func GetVersionsByServiceID(db *gorm.DB, c *gin.Context) {

	var uri request.ServiceURI
	if err := c.ShouldBindUri(&uri); err != nil {
		response.BadRequest(c, "Missing ID", err.Error())
		return
	}

	repo := repository.NewVersionRepository(db)

	serviceHndlr := serviceHandler.NewVersion(repo)

	// Call the service layer to retrieve the versions  by Service ID
	versions, err := serviceHndlr.GetVersionsByServiceID(uri.ServiceID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NotFound(c, "Service not found")
		} else {
			response.InternalServerError(c, "Failed to retrieve service", err.Error())
		}
		return
	}

	response.OK(c, versions)
}
