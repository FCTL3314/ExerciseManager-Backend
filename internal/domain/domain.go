package domain

type FilterParams struct {
	Limit  int
	Offset int
	Query  interface{}
	Args   []interface{}
}

type OrderParams struct {
	Order string
}

type Params struct {
	FilterParams
	OrderParams
}

type Getter[T any] interface {
	Get(query interface{}, args ...interface{}) (*T, error)
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
