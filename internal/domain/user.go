package domain

import (
	"time"
)

type ToUserConverter interface {
	ToUser() *User
}

type ResponseUser struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type UserBase struct {
	Username string `json:"username" validate:"required,min=4,max=16"`
}

type CreateUser struct {
	UserBase
	Password string `json:"password" validate:"required,min=6,max=128"`
}

func (cu *CreateUser) ToUser() *User {
	return &User{
		Username: cu.Username,
		Password: cu.Password,
	}
}

type UpdateUser struct {
	UserBase
}

func (uu *UpdateUser) ToUser() *User {
	return &User{
		Username: uu.Username,
	}
}

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
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
