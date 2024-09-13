package domain

type GetRouter interface {
	RegisterGet()
}

type ListRouter interface {
	RegisterList()
}

type CreateRouter interface {
	RegisterCreate()
}

type UpdateRouter interface {
	RegisterUpdate()
}

type DeleteRouter interface {
	RegisterDelete()
}

type AllRouter interface {
	RegisterAll()
}

type Router interface {
	GetRouter
	ListRouter
	CreateRouter
	UpdateRouter
	DeleteRouter
	AllRouter
}
