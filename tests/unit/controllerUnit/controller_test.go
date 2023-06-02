package controllerUnit

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/siparisa/ServiceCatalog/internal/controller"
	"github.com/siparisa/ServiceCatalog/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var db *gorm.DB

type Service struct {
	gorm.Model
	Name        *string
	Description *string
	Versions    []Version
}

type Version struct {
	gorm.Model
	ServiceID uint
	Version   string
}

func TestMain(m *testing.M) {
	// Create an in-memory SQLite database connection using GORM
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to open database connection: " + err.Error())
	}

	err = db.AutoMigrate(&Service{}, &Version{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	exitCode := m.Run()

	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database connection: " + err.Error())
	}
	err = sqlDB.Close()
	if err != nil {
		panic("failed to close database connection: " + err.Error())
	}

	os.Exit(exitCode)
}

func TestGetServices(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/services?name=Service1&description=Description1", nil)

	q := req.URL.Query()
	q.Add("name", "Service1")
	q.Add("description", "Description1")
	req.URL.RawQuery = q.Encode()

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

func TestGetServiceByID(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/services/1", nil)

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Set the path parameter for the request
	c.Params = append(c.Params, gin.Param{
		Key:   "serviceID",
		Value: "1",
	})

	controller.GetServiceByID(db, c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []entity.Service
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	name := "Service1"
	desc := "Description1"
	expectedResponse := entity.Service{
		Model:       gorm.Model{ID: 1},
		Name:        &name,
		Description: &desc,
		Versions: []entity.Version{
			{
				Model:     gorm.Model{ID: 1},
				ServiceID: 1,
				Version:   "1.0",
			},
			{
				Model:     gorm.Model{ID: 2},
				ServiceID: 1,
				Version:   "2.0",
			},
		},
	}

	assert.Equal(t, expectedResponse, response)
}

func TestGetServiceByID_InvalidID(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/services/abc", nil)

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	c.Params = append(c.Params, gin.Param{
		Key:   "serviceID",
		Value: "abc",
	})

	controller.GetServiceByID(db, c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetServices_Error(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/services", nil)

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	q := req.URL.Query()
	q.Add("invalidParam", "value")
	req.URL.RawQuery = q.Encode()

	controller.GetServices(db, c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
