package domain

type Getter[T any] interface {
	Get(query interface{}, args ...interface{}) (*T, error)
}

type Lister[T any] interface {
	List(query interface{}, args ...interface{}) ([]*T, error)
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
