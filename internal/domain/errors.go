package domain

import (
	"fmt"
)

type PaginationLimitExceededError struct {
	MaxLimit int
}

func (e *PaginationLimitExceededError) Error() string {
	return fmt.Sprintf("Pagination limit exceeded, maximum allowed limit is %d", e.MaxLimit)
}

type ObjectUniqueConstraintError struct {
	Field string
}

func (e *ObjectUniqueConstraintError) Error() string {
	return fmt.Sprintf(
		"Unique constraint violation on field \"%s\". An object with this value already exists.",
		e.Field,
	)
}
