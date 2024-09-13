package domain

type Getter[T any] interface {
	Get(params *FilterParams) (*T, error)
}

type Lister[T any] interface {
	List(params *Params) ([]*T, error)
}

type Creator[T any] interface {
	Create(entity *T) (*T, error)
}

type Updater[T any] interface {
	Update(entity *T) (*T, error)
}

type Deleter[T any] interface {
	Delete(id uint) error
}
