package domain

type FilterParams struct {
	Query interface{}
	Args  []interface{}
}

type PaginationParams struct {
	Limit  int
	Offset int
}

type OrderParams struct {
	Order string
}

type Params struct {
	Filter     FilterParams
	Pagination PaginationParams
	OrderParams
}

type Getter[T any] interface {
	Get(params *FilterParams) (*T, error)
}

type Lister[T any] interface {
	List(params *Params) ([]*T, error)
}

type Creator[T any] interface {
	Create(entity *T) (*T, error)
}

type Deleter[T any] interface {
	Delete(id uint) error
}

type Repository[T any] interface {
	Getter[T]
	Creator[T]
	Lister[T]
	Deleter[T]
}

type Usecase[T any] interface {
	Getter[T]
	Creator[T]
	Lister[T]
	Deleter[T]
}

type Controller[T any] interface {
	Getter[T]
	Creator[T]
	Lister[T]
	Deleter[T]
}
