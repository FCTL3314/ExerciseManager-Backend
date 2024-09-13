package domain

import "fmt"

type PaginationLimitExceededError struct {
	MaxLimit int
}

func (e *PaginationLimitExceededError) Error() string {
	return fmt.Sprintf("Pagination limit exceeded, maximum allowed limit is %d", e.MaxLimit)
}
