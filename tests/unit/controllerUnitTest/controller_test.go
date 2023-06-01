package controllerUnitTest

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/siparisa/ServiceCatalog/internal/controller"
	"github.com/siparisa/ServiceCatalog/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetServices(t *testing.T) {
	// Create a new HTTP request and response recorder
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/services?name=Service1&description=Description1", nil)

	// Set the query parameters for the request
	q := req.URL.Query()
	q.Add("name", "Service1")
	q.Add("description", "Description1")
	req.URL.RawQuery = q.Encode()

	// Create a new Gin context for the request
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	controller.GetServices(db, c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []entity.Service
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	name1 := "Service1"
	name2 := "Service2"
	desc1 := "Description1"
	desc2 := "Description2"
	// Assert the response content
	expectedResponse := []entity.Service{
		{
			Model:       gorm.Model{ID: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}},
			Name:        &name1,
			Description: &desc1,
			Versions: []entity.Version{
				{
					Model:     gorm.Model{ID: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}},
					ServiceID: 1,
					Version:   "1.0",
				},
				{
					Model:     gorm.Model{ID: 2, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}},
					ServiceID: 1,
					Version:   "2.0",
				},
			},
		},
		{
			Model:       gorm.Model{ID: 2, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}},
			Name:        &name2,
			Description: &desc2,
			Versions: []entity.Version{
				{
					Model:     gorm.Model{ID: 3, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}},
					ServiceID: 2,
					Version:   "1.0",
				},
			},
		},
	}

	assert.Equal(t, expectedResponse, response)

}

//func TestGetServiceByID(t *testing.T) {
//
//	// Create a new HTTP request and response recorder
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("GET", "/services/1", nil)
//
//	// Create a new Gin context for the request
//	c, _ := gin.CreateTestContext(w)
//	c.Request = req
//
//	// Set the path parameter for the request
//	c.Params = append(c.Params, gin.Param{
//		Key:   "serviceID",
//		Value: "1",
//	})
//
//	controller.GetServiceByID(db, c)
//
//	assert.Equal(t, http.StatusOK, w.Code)
//
//	var response []entity.Service
//	_ = json.Unmarshal(w.Body.Bytes(), &response)
//
//	name := "Service1"
//	desc := "Description1"
//	expectedResponse := entity.Service{
//		Model:       gorm.Model{ID: 1},
//		Name:        &name,
//		Description: &desc,
//		Versions: []entity.Version{
//			{
//				Model:     gorm.Model{ID: 1},
//				ServiceID: 1,
//				Version:   "1.0",
//			},
//			{
//				Model:     gorm.Model{ID: 2},
//				ServiceID: 1,
//				Version:   "2.0",
//			},
//		},
//	}
//
//	assert.Equal(t, expectedResponse, response)
//}
//
//func TestGetServiceByID_InvalidID(t *testing.T) {
//
//	// Create a new HTTP request and response recorder
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("GET", "/services/abc", nil)
//
//	// Create a new Gin context for the request
//	c, _ := gin.CreateTestContext(w)
//	c.Request = req
//
//	// Set the path parameter for the request
//	c.Params = append(c.Params, gin.Param{
//		Key:   "serviceID",
//		Value: "abc",
//	})
//
//	controller.GetServiceByID(db, c)
//
//	assert.Equal(t, http.StatusBadRequest, w.Code)
//}
//
//func TestGetServices_Error(t *testing.T) {
//
//	// Create a new HTTP request and response recorder
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("GET", "/services", nil)
//
//	// Create a new Gin context for the request
//	c, _ := gin.CreateTestContext(w)
//	c.Request = req
//
//	// Simulate invalid query parameters
//	// You can modify this based on the specific validation rules of your application
//	q := req.URL.Query()
//	q.Add("invalidParam", "value")
//	req.URL.RawQuery = q.Encode()
//
//	controller.GetServices(db, c)
//
//	assert.Equal(t, http.StatusBadRequest, w.Code)
//}
