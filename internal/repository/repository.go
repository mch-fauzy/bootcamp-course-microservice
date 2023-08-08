package repository

import "bootcamp-course-microservice/infras"

type Repository interface {
	CourseRepository
}

type RepositoryImpl struct {
	DB *infras.Conn
}

func ProvideRepo(db *infras.Conn) *RepositoryImpl {
	return &RepositoryImpl{
		DB: db,
	}
}
