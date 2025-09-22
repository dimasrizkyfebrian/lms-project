// internal/repository/user_repository.go
package repository

import (
	"lms-project/backend/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Create inserts a new user record into the database.
func (r *UserRepository) Create(user *model.User) error {
	// The Create function from GORM will handle the SQL INSERT statement.
	result := r.DB.Create(user)
	return result.Error
}