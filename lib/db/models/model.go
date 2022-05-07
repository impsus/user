package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string   `gorm:"primaryKey"`
	Keys []ApiKey `gorm:"foreignKey:UserId"`
}

type ApiKey struct {
	gorm.Model
	UserId      string
	Id          uint
	Description string
	Key         string
}
