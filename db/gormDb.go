package db

import (
	"coursesApi/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGormDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("oops, something goes wrong: %s", err.Error()))
	}

	db.AutoMigrate(&models.Course{})

	return db
}
