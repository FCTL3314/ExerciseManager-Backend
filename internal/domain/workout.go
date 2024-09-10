package domain

import (
	"time"
)

type Workout struct {
	ID          uint
	Name        string
	Description string
	UserID      uint
	User        User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
