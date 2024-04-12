package services

import (
	"coursesApi/repo"
)

type CoursesService struct {
	repo *repo.GormCoursesRepo
}

func NewCoursesService(r *repo.GormCoursesRepo) *CoursesService {
	return &CoursesService{
		repo: r,
	}
}

// func (s *CoursesService) CreateIfNotExist() (bool, error) {
// }
