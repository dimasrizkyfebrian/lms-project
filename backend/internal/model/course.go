// internal/model/course.go
package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Course defines the course model for the database.
type Course struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Title       string         `gorm:"not null"`
	Description string
	UserID      uint
}