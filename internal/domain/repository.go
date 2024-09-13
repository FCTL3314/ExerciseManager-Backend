package domain

type Getter[T any] interface {
	Get(params *FilterParams) (*T, error)
}

type Fetcher[T any] interface {
	Fetch(params *Params) ([]*T, error)
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

type Counter[T any] interface {
	Count() (int64, error)
}

type Repository[T any] interface {
	Getter[T]
	Fetcher[T]
	Creator[T]
	Updater[T]
	Deleter[T]
	Counter[T]
}

type UserRepository interface {
	Repository[User]
	GetByUsername(username string) (*User, error)
}
