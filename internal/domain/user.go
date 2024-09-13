package domain

import (
	"time"
)

type ResponseUser struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID        uint
	Username  string
	Password  string
	CreatedAt time.Time
}

func (u *User) ToResponseUser() *ResponseUser {
	return &ResponseUser{
		ID:        u.ID,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}
}

func ToResponseUsers(users []*User) []*ResponseUser {
	responseUsers := make([]*ResponseUser, len(users))
	for i, user := range users {
		responseUsers[i] = user.ToResponseUser()
	}
	return responseUsers
}
