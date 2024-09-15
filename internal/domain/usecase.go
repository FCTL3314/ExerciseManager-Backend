package domain

type UserUsecase interface {
	GetById(id uint) (*User, error)
	Get(params *FilterParams) (*User, error)
	List(params *Params) (*PaginatedResult[*User], error)
	Create(createUser *CreateUserRequest) (*User, error)
	Login(loginUser *LoginUserRequest) (*TokensResponse, error)
	RefreshTokens(refreshTokenRequest *RefreshTokenRequest) (*TokensResponse, error)
	Update(authUserId uint, id uint, updateUser *UpdateUserRequest) (*User, error)
	Delete(authUserId uint, id uint) error
}
