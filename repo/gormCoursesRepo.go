package repo

import (
	"coursesApi/models"
	"fmt"

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
	f := r.db.Table("teeest1").Find(course)
	fmt.Println(f)
	return course, tx.Error
}

func (r *GormCoursesRepo) Update(course *models.Course, id int) error {
	tx := r.db.Model(course).Where(id).Updates(course)

	return tx.Error
}

func (r *GormCoursesRepo) Delete(id int) error {
	course := &models.Course{}
	tx := r.db.Delete(course, id)

	return tx.Error
}

func (r *GormCoursesRepo) GetList(limit int, offset int) (*[]models.Course, error) {
	courses := &[]models.Course{}
	tx := r.db.Limit(limit).Offset(offset).Find(courses)

	return courses, tx.Error
}
