// internal/model/user.go
package model

import "gorm.io/gorm"

// User defines the user model for the database.
type User struct {
	gorm.Model // Includes fields like ID, CreatedAt, UpdatedAt, DeletedAt
	Username   string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
}