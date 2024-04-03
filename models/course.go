package models

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Description   string
	Name          string
	Author        string
	StartTime     time.Time
	NumberOfSeats int
}
