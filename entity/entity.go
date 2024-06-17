package entity

import "gorm.io/gorm"

type Entity interface {
	Initialize(dbConnection *gorm.DB)
}
