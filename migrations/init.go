package migrations

import "gorm.io/gorm"

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
