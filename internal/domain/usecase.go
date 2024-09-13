package domain

type UserUsecase interface {
	Get(params *FilterParams) (*User, error)
	List(params *Params) (*PaginatedResult[*User], error)
	Create(user *User) (*User, error)
	Update(user *User) (*User, error)
	Delete(id uint) error
}
