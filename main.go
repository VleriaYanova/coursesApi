package main

import (
	"coursesApi/db"
	"coursesApi/models"
	"coursesApi/repo"
	"time"
)

func main() {
	db := db.NewGormDb()

	repository := repo.NewGormCoursesRepo(db)

	c, err := repository.Create(&models.Course{
		Author:        "Vasya",
		Description:   "pispis",
		StartTime:     time.Now(),
		NumberOfSeats: 13,
	})

}
