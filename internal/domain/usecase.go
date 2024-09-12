package domain

type Usecase[T any] interface {
	Getter[T]
	Creator[T]
	Lister[T]
	Deleter[T]
}

type UserUsecase interface {
	Usecase[User]
}
