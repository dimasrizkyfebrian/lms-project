// internal/service/user_service.go
package service

import (
	"lms-project/backend/internal/model"
	"lms-project/backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// Register handles the business logic for user registration.
func (s *UserService) Register(username, password string) (*model.User, error) {
	// Hash the password using bcrypt for security.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create a new User model instance.
	newUser := &model.User{
		Username: username,
		Password: string(hashedPassword),
	}

	// Call the repository to save the user to the database.
	err = s.userRepo.Create(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}