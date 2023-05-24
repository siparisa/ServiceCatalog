package tests

import (
	"github.com/siparisa/ServiceCatalog/internal/entity"
	"gorm.io/gorm"
)

// MigrateDB is a Helper function to perform database migrations
func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&entity.Service{}, &entity.Version{})
}

func SetupTestData(db *gorm.DB) {
	db.Exec("DELETE FROM services")
	db.Exec("DELETE FROM versions")

	name1 := "Service1"
	// Create test services
	service1 := entity.Service{
		Name:        &name1,
		Description: "Description1",
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
	service2 := entity.Service{
		Name:        &name2,
		Description: "Description2",
		Versions: []entity.Version{
			{
				Version: "v1",
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
