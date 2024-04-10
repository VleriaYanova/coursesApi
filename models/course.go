package models

import (
	"time"
)

type Course struct {
	ID            uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Description   string
	Name          string
	Author        string
	StartTime     time.Time
	NumberOfSeats int
}
