// internal/service/course_service.go
package service

import (
	"lms-project/backend/internal/model"
	"lms-project/backend/internal/repository"
	"strconv"

	"github.com/google/uuid"
)

type CourseService struct {
	courseRepo *repository.CourseRepository
}

func NewCourseService(courseRepo *repository.CourseRepository) *CourseService {
	return &CourseService{courseRepo: courseRepo}
}

func (s *CourseService) GetAll() ([]model.Course, error) {
	return s.courseRepo.FindAll()
}

func (s *CourseService) CreateCourse(title, description, userIDStr string) (*model.Course, error) {
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return nil, err
	}

	newCourse := &model.Course{
		Title:       title,
		Description: description,
		UserID:      uint(userID),
	}

	err = s.courseRepo.Create(newCourse)
	if err != nil {
		return nil, err
	}

	return newCourse, nil
}

// GetCourseByID now parses the string into a uuid.UUID
func (s *CourseService) GetCourseByID(idStr string) (*model.Course, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	return s.courseRepo.FindByID(id)
}