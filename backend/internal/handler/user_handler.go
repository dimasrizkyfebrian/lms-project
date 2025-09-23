// internal/handler/user_handler.go
package handler

import (
	"lms-project/backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler now depends on the UserService.
type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// A struct to bind the JSON request body for registration.
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register validates the input, calls the service, and sends a response.
func (h *UserHandler) Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Register(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to register user. Username might already exist."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"userId":  user.ID,
	})
}

// -- FUNGSI BARU DI BAWAH INI --

// LoginInput defines the structure for the login request body.
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login handles the user login request.
func (h *UserHandler) Login(c *gin.Context) {
	var input LoginInput

	// Bind the incoming JSON to the input struct.
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the user service to perform the login logic.
	token, err := h.userService.Login(input.Username, input.Password)
	if err != nil {
		// If the service returns an error (e.g., "invalid credentials").
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Send a successful response with the JWT token.
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GetMe is the handler for the protected /me route.
func (h *UserHandler) GetMe(c *gin.Context) {
	// Retrieve the userID set by the AuthMiddleware.
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Call the service to get the user's profile.
	user, err := h.userService.GetUserProfile(userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return user data (excluding the password).
	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
	})
}