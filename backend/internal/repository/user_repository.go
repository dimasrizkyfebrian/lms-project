// internal/repository/user_repository.go
package repository

import (
	"lms-project/backend/internal/model"
	"strconv"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *model.User) error {
	result := r.DB.Create(user)
	return result.Error
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	result := r.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}


// FindByID retrieves a single user by their primary key (ID).
func (r *UserRepository) FindByID(id string) (*model.User, error) {
	var user model.User
	// Convert the string ID to an integer for GORM.
	userID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	result := r.DB.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}