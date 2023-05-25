package entity

import "gorm.io/gorm"

// Version represents a version entity in the system.
type Version struct {
	gorm.Model // Embedded struct from GORM providing common fields like ID, CreatedAt, UpdatedAt, DeletedAt
	ServiceID  uint
	Version    string
}

// TableName returns the name of the database table associated with the Version struct.
func (Version) TableName() string {
	return "versions"
}
