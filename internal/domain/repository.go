package domain

type Repository[T any] interface {
	Getter[T]
	Creator[T]
	Updater[T]
	Lister[T]
	Deleter[T]
}

type UserRepository interface {
	Repository[User]
	GetByUsername(username string) (*User, error)
}
