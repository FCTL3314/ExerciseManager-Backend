package domain

type Getter[T any] interface {
	Get(conds ...interface{}) (*T, error)
}

type Creator[T any] interface {
	Create(entity *T) error
}

type Lister[T any] interface {
	List(conds ...interface{}) ([]T, error)
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
