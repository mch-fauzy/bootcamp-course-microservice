package repository

import (
	"bootcamp-course-microservice/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type CourseRepository interface {
	ReadCourseByUserID(userID string) ([]models.Course, error)
	CreateCourse(course *models.Course) error
}

func (r *RepositoryImpl) ReadCourseByUserID(userID string) ([]models.Course, error) {
	query := "SELECT * FROM bootcamp_courses WHERE user_id = ?"

	var course []models.Course
	err := r.DB.Read.Select(&course, query, userID)
	if err != nil {
		log.Error().Err(err).Msg("Something went wrong")
		return nil, err
	}
	return course, nil
}

func (r *RepositoryImpl) CreateCourse(course *models.Course) error {
	query :=
		`
	INSERT INTO bootcamp_courses (id, user_id, title, content, created_at, updated_at) 
	VALUES (?, ?, ?, ?, ?, ?)
	`
	course.ID = uuid.New().String()
	course.CreatedAt = time.Now()
	course.UpdatedAt = time.Now()
	_, err := r.DB.Write.Exec(
		query,
		course.ID,
		course.UserID,
		course.Title,
		course.Content,
		course.CreatedAt,
		course.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
