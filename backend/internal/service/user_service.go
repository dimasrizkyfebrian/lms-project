// internal/service/user_service.go
package service

import (
	"errors"
	"lms-project/backend/internal/model"
	"lms-project/backend/internal/repository"
	"strconv" // Import strconv for correct integer to string conversion
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// We removed the global jwtKey variable from here.

type UserService struct {
	userRepo *repository.UserRepository
	jwtKey   []byte // The JWT secret key is now part of the service struct
}

// NewUserService now accepts the JWT secret key as an argument.
func NewUserService(userRepo *repository.UserRepository, jwtKey string) *UserService {
	return &UserService{
		userRepo: userRepo,
		jwtKey:   []byte(jwtKey),
	}
}

// ... Register function remains the same ...
func (s *UserService) Register(username, password string) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		Username: username,
		Password: string(hashedPassword),
	}
	err = s.userRepo.Create(newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}


func (s *UserService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid credentials")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		// FIX: Correctly convert user ID (uint) to string for the JWT subject.
		Subject:   strconv.FormatUint(uint64(user.ID), 10),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Use the jwtKey from the struct, not a global variable.
	tokenString, err := token.SignedString(s.jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ... GetUserProfile remains the same ...
func (s *UserService) GetUserProfile(id string) (*model.User, error) {
	return s.userRepo.FindByID(id)
}