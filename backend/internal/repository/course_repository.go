// internal/repository/course_repository.go
package repository

import (
	"lms-project/backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseRepository struct {
	DB *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{DB: db}
}

func (r *CourseRepository) FindAll() ([]model.Course, error) {
	var courses []model.Course
	result := r.DB.Find(&courses)
	return courses, result.Error
}

func (r *CourseRepository) Create(course *model.Course) error {
	result := r.DB.Create(course)
	return result.Error
}

// FindByID now accepts a uuid.UUID
func (r *CourseRepository) FindByID(id uuid.UUID) (*model.Course, error) {
	var course model.Course
	result := r.DB.First(&course, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &course, nil
}