package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	name string
	keys []ApiKey
}

type ApiKey struct {
	description string
	key         string
}
