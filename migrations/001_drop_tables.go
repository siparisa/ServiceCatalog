package migrations

import (
	"gorm.io/gorm"
)

// RollbackServicesTable drops services and versions tables
func RollbackServicesTable(db *gorm.DB) error {
	err := db.Migrator().DropTable(&Service{}, &Version{})
	if err != nil {
		return err
	}

	return nil
}
