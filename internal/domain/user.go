package domain

import "time"

type User struct {
	ID        int64
	Username  string
	Password  string
	CreatedAt time.Time
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=4,max=16"`
	Password string `json:"password" binding:"required,min=6,max=128"`
}

type UpdateUserRequest struct {
	Username *string `json:"username,omitempty" binding:"omitempty,min=4,max=16"`
}

type LoginUserRequest struct {
	Username string `json:"username" binding:"required,min=4,max=16"`
	Password string `json:"password" binding:"required,min=6,max=128"`
}

type ResponseUser struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type TokensResponse struct {
	AccessToken  string `json:"access_token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func NewUserFromCreateRequest(req *CreateUserRequest) *User {
	return &User{
		Username: req.Username,
		Password: req.Password, // Password must be cached
	}
}

func (u *User) ToResponseUser() *ResponseUser {
	return &ResponseUser{
		ID:        u.ID,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}
}

func (u *User) ApplyUpdate(req *UpdateUserRequest) {
	if req.Username != nil {
		u.Username = *req.Username
	}
}

func ToResponseUsers(users []*User) []*ResponseUser {
	responseUsers := make([]*ResponseUser, len(users))
	for i, user := range users {
		responseUsers[i] = user.ToResponseUser()
	}
	return responseUsers
}
