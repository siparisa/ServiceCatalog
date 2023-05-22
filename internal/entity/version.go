package entity

import "gorm.io/gorm"

type Version struct {
	gorm.Model
	ServiceID uint
	Version   string
}

func (Version) TableName() string {
	return "versions"
}
