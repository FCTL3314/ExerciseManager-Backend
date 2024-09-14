package domain

type UserUsecase interface {
	GetById(id uint) (*User, error)
	Get(params *FilterParams) (*User, error)
	List(params *Params) (*PaginatedResult[*User], error)
	Create(createUser *CreateUser) (*User, error)
	Login(loginUser *LoginUser) (*SuccessLoginResponse, error)
	Update(authUserId uint, id uint, updateUser *UpdateUser) (*User, error)
	Delete(authUserId uint, id uint) error
}
