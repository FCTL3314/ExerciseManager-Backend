package domain

type Usecase[T any] interface {
	Getter[T]
	Lister[T]
	Creator[T]
	Updater[T]
	Deleter[T]
}

type UserUsecase interface {
	Usecase[User]
}
