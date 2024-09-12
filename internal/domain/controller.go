package domain

type Controller[T any] interface {
	Getter[T]
	Creator[T]
	Lister[T]
	Deleter[T]
}

type UserController interface {
	Controller[User]
}
