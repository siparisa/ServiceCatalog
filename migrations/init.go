package migrations

import (
	"gorm.io/gorm"
)

type Service struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   int64
	UpdatedAt   int64
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	Description string
}

func (Service) TableName() string {
	return "services"
}

type Version struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt int64
	UpdatedAt int64
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ServiceID uint
	Version   string
}

func (Version) TableName() string {
	return "versions"
}

//package migrations
//
//import "gorm.io/gorm"
//
//type Service struct {
//	gorm.Model
//	Name        string
//	Description string
//}
//
//func (Service) TableName() string {
//	return "services"
//}
//
//type Version struct {
//	gorm.Model
//	ServiceID uint
//	Version   string
//}
//
//func (Version) TableName() string {
//	return "versions"
//}
