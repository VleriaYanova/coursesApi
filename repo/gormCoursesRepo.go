package repo

import (
	"coursesApi/models"

	"gorm.io/gorm"
)

type GormCoursesRepo struct {
	db *gorm.DB
}

func NewGormCoursesRepo(db *gorm.DB) *GormCoursesRepo {
	return &GormCoursesRepo{db: db}
}

func (r *GormCoursesRepo) Create(course *models.Course) (*models.Course, error) {
	tx := r.db.Create(course)

	return course, tx.Error
}

func (r *GormCoursesRepo) Get() {

}

func (r *GormCoursesRepo) Update() {}

func (r *GormCoursesRepo) Delete() {}
