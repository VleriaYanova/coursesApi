package models

import (
	"time"
)

type Test struct {
	ID            uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Description   string
	Name          string
	NumberOfSeats int
}
