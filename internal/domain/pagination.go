package domain

const (
	MaxUserPaginationLimit     = 32
	MaxWorkoutPaginationLimit  = 32
	MaxExercisePaginationLimit = 64
)

type PaginatedResult[T any] struct {
	Results []T
	Count   int64
}
