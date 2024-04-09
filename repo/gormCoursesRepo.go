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

func (r *GormCoursesRepo) Get(id int) (*models.Course, error) {
	course := &models.Course{}
	tx := r.db.First(course, id)

	return course, tx.Error
}

func (r *GormCoursesRepo) Update(UpdCourse *models.Course, id int) error {
	tx := r.db.Model(UpdCourse).Where(id).Updates(UpdCourse)

	return tx.Error
}

func (r *GormCoursesRepo) Delete(id int) error {
	course := &models.Course{}
	tx := r.db.Delete(course, id)

	return tx.Error
}
