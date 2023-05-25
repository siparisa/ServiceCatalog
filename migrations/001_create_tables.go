package migrations

import (
	"gorm.io/gorm"
)

// MigrateServicesTable creates services and versions tables
func MigrateServicesTable(db *gorm.DB) error {
	err := db.AutoMigrate(&Service{}, &Version{})
	if err != nil {
		return err
	}

	return nil
}
