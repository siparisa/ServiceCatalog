package controllerUnitTest

import (
	"github.com/siparisa/ServiceCatalog/internal/entity"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	var err error
	// Create an in-memory SQLite database connection using GORM
	db, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to open database connection: " + err.Error())
	}

	// Perform database migrations and setup test data
	migrateDB(db)
	setupTestData(db)

	// Run the tests
	exitCode := m.Run()

	sqlDB, err := db.DB()
	defer sqlDB.Close()

	if err != nil {
		panic("failed to close database connection: " + err.Error())
	}

	os.Exit(exitCode)
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&entity.Service{}, &entity.Version{})
}

func setupTestData(db *gorm.DB) {
	db.Exec("DELETE FROM services")
	db.Exec("DELETE FROM versions")

	name1 := "Service1"
	desc1 := "Description1"
	// Create test services
	service1 := entity.Service{
		Name:        &name1,
		Description: &desc1,
		Versions: []entity.Version{
			{
				Version: "1.0",
			},
			{
				Version: "2.0",
			},
		},
	}
	db.Save(&service1)

	// Insert versions
	for _, v := range service1.Versions {
		version := entity.Version{
			ServiceID: service1.ID,
			Version:   v.Version,
		}
		db.Create(&version)
	}

	name2 := "Service2"
	desc2 := "Description2"
	service2 := entity.Service{
		Name:        &name2,
		Description: &desc2,
		Versions: []entity.Version{
			{
				Version: "1.0",
			},
		},
	}

	db.Create(&service2)
	// Insert versions
	for _, v := range service2.Versions {
		version := entity.Version{
			ServiceID: service2.ID,
			Version:   v.Version,
		}
		db.Create(&version)
	}
}
