package domain

import (
	"time"
)

type Workout struct {
	ID          uint
	Name        string
	Description string
	UserID      uint
	User        *User
	Exercises   []*Exercise `gorm:"many2many:workout_exercises;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
