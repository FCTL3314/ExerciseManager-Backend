package domain

import (
	"time"
)

type User struct {
	ID        uint
	Username  string
	Password  string
	CreatedAt time.Time
}

type UserRepository interface {
	Repository[User]
	GetByUsername(username string) (User, error)
}
