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
	Username string `json:"username" binding:"required,min=4,max=16"`
}

func (ub *UserBase) ToUser() *User {
	return &User{
		Username: ub.Username,
	}
}

func (ub *UserBase) ApplyToUser(user *User) *User {
	user.Username = ub.Username
	return user
}

func (cu *CreateUserRequest) ToUser() *User {
	user := cu.UserBase.ToUser()
	user.Password = cu.Password
	return user
}

type CreateUserRequest struct {
	UserBase
	Password string `json:"password" binding:"required,min=6,max=128"`
}

type UpdateUserRequest struct {
	UserBase
}

type User struct {
	Id        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) ToResponseUser() *ResponseUser {
	return &ResponseUser{
		ID:        u.Id,
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

type LoginUserRequest struct {
	UserBase
	Password string `json:"password" binding:"required,min=6,max=128"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type TokensResponse struct {
	AccessToken  string `json:"access_token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}
