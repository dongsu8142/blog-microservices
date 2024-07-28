package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string
	Password string `gorm:"not null"`
}
