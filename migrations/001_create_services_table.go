package migrations

import (
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name        string
	Description string
}

func (Service) TableName() string {
	return "services"
}

type Version struct {
	gorm.Model
	ServiceID uint
	Version   string
}

func (Version) TableName() string {
	return "versions"
}

func MigrateServicesTable(db *gorm.DB) error {
	err := db.AutoMigrate(&Service{}, &Version{})
	if err != nil {
		return err
	}

	return nil
}

func RollbackServicesTable(db *gorm.DB) error {
	err := db.Migrator().DropTable(&Service{}, &Version{})
	if err != nil {
		return err
	}

	return nil
}
