package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Keys []ApiKey
}

type ApiKey struct {
	Description string
	Key         string
}
