// internal/handler/course_handler.go
package handler

import (
	"lms-project/backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	courseService *service.CourseService
}

func NewCourseHandler(courseService *service.CourseService) *CourseHandler {
	return &CourseHandler{courseService: courseService}
}

// ... GetAll and CreateCourseInput ...
func (h *CourseHandler) GetAll(c *gin.Context) {
	courses, err := h.courseService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve courses"})
		return
	}

	c.JSON(http.StatusOK, courses)
}
type CreateCourseInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}


// ... Create function ...
func (h *CourseHandler) Create(c *gin.Context) {
	var input CreateCourseInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	course, err := h.courseService.CreateCourse(input.Title, input.Description, userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}

	c.JSON(http.StatusCreated, course)
}


// GetByID handles the request to get a single course by its ID.
func (h *CourseHandler) GetByID(c *gin.Context) {
	// Get the ID from the URL parameter.
	id := c.Param("id")

	course, err := h.courseService.GetCourseByID(id)
	if err != nil {
		// If the course is not found, or the ID is invalid.
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, course)
}