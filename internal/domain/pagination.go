package domain

type PaginatedResult[T any] struct {
	Results []T
	Count   int64
}
