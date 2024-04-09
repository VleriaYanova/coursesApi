package main

import (
	"coursesApi/db"
	"coursesApi/models"
	"coursesApi/repo"
	"fmt"
	// "time"
)

func main() {
	db := db.NewGormDb()

	repository := repo.NewGormCoursesRepo(db)

	c := repository.Update(&models.Course{Author: "AAAA"}, 1)

	fmt.Println(c)

}
