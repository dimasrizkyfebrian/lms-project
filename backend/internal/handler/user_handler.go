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

	// Bind the incoming JSON to the input struct.
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the user service to perform the registration logic.
	user, err := h.userService.Register(input.Username, input.Password)
	if err != nil {
		// This could be a duplicate username error or something else.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to register user. Username might already exist."})
		return
	}

	// Send a successful response.
	// It's good practice to not return the password hash.
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"userId":  user.ID,
	})
}