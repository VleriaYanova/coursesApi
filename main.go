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

	c := repository.Delete(3)

	fmt.Println(c)

}
