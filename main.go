package main

import (
	"coursesApi/db"
	"coursesApi/repo"
	"fmt"
	// "time"
)

func main() {
	db := db.NewGormDb()

	repository := repo.NewGormCoursesRepo(db)

	// c, err := repository.Create(&models.Course{Name: "Vasya", Author: "ls;dgh"})

	c, err := repository.GetList(5)

	fmt.Println(c, err)

}
