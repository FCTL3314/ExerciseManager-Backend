package domain

type GetterById[T any] interface {
	GetById(id int64) (*T, error)
}

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
	Delete(id int64) error
}

type Counter[T any] interface {
	Count(params *FilterParams) (int64, error)
}

type Repository[T any] interface {
	GetterById[T]
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

type WorkoutRepository interface {
	Repository[Workout]
}
