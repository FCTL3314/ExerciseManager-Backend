package domain

const (
	MaxUserPaginationLimit    = 32
	MaxWorkoutPaginationLimit = 32
)

type PaginatedResult[T any] struct {
	Results []T
	Count   int64
}
