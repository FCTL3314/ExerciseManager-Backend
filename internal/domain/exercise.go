package domain

import (
	"time"
)

type Exercise struct {
	ID          uint
	Name        string
	Description string
	Duration    time.Duration
}
