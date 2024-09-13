package domain

import (
	"time"
)

type ResponseUser struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (cr *CreateUser) ToUser() *User {
	return &User{
		Username: cr.Username,
		Password: cr.Password,
	}
}

type User struct {
	ID        uint      `json:"id" binding:"-"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"-"`
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
