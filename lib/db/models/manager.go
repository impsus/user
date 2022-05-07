package models

import "gorm.io/gorm"

func RegisterModels(db *gorm.DB) {
	db.AutoMigrate(&ApiKey{})
	db.AutoMigrate(&User{})
}
