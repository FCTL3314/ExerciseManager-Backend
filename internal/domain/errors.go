package domain

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrObjectNotFound         = errors.New("object not found")
	ErrAccessDenied           = errors.New("access denied")
	ErrInvalidAuthCredentials = errors.New("invalid auth credentials")
)

type ErrPaginationLimitExceeded struct {
	MaxLimit int
}

func (e *ErrPaginationLimitExceeded) Error() string {
	return fmt.Sprintf("Pagination limit exceeded, maximum allowed limit is %d", e.MaxLimit)
}

type ErrObjectUniqueConstraint struct {
	Fields []string
}

func (e *ErrObjectUniqueConstraint) Error() string {
	return fmt.Sprintf(
		"Unique constraint violation on fields \"%s\". An object with values for these fields already exists.",
		strings.Join(e.Fields, ","),
	)
}
