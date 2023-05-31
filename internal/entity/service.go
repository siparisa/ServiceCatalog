package entity

import "gorm.io/gorm"

// Service represents a service entity in the system.
type Service struct {
	gorm.Model  // Embedded struct from GORM providing common fields like ID, CreatedAt, UpdatedAt, DeletedAt
	Name        *string
	Description *string
	Versions    []Version
}

// TableName returns the name of the database table associated with the Service struct.
func (Service) TableName() string {
	return "services"
}
