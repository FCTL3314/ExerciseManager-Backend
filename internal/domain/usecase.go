package domain

type UserUsecase interface {
	GetById(id int64) (*User, error)
	Get(params *FilterParams) (*User, error)
	List(params *Params) (*PaginatedResult[*User], error)
	Create(createUser *CreateUserRequest) (*User, error)
	Login(loginUser *LoginUserRequest) (*TokensResponse, error)
	RefreshTokens(refreshTokenRequest *RefreshTokenRequest) (*TokensResponse, error)
	Update(authUserId int64, id int64, updateUser *UpdateUserRequest) (*User, error)
	Delete(authUserId int64, id int64) error
}
