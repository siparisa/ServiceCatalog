package entity

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name        *string
	Description string
	Versions    []Version
}

func (Service) TableName() string {
	return "services"
}
