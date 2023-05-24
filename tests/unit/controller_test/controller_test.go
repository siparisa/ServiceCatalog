package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/siparisa/ServiceCatalog/internal/controller"
	"github.com/siparisa/ServiceCatalog/internal/entity"
	"github.com/siparisa/ServiceCatalog/tests"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetServices(t *testing.T) {

	// Create an in-memory SQLite database connection using GORM
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	// Perform database migrations and setup test data
	tests.MigrateDB(db)
	tests.SetupTestData(db)

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

	name1 := "Service 1"
	name2 := "Service 2"

	// Assert the response content
	expectedResponse := []entity.Service{
		{
			Model:       gorm.Model{ID: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}},
			Name:        &name1,
			Description: "Description 1",
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
			Description: "Description 2",
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
