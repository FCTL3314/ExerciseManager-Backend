package controller

import (
	"ExerciseManager/internal/domain"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func tryToHandleErr(c *gin.Context, err error) (IsHandled bool) {
	if errors.Is(err, domain.ErrObjectNotFound) {
		c.JSON(http.StatusNotFound, domain.NotFoundResponse)
		return true
	}

	if errors.Is(err, domain.ErrAccessDenied) {
		c.JSON(http.StatusForbidden, domain.ForbiddenResponse)
		return true
	}

	if errors.Is(err, domain.ErrInvalidAuthCredentials) {
		c.JSON(http.StatusUnauthorized, domain.InvalidAuthCredentialsResponse)
		return true
	}

	var limitErr *domain.ErrPaginationLimitExceeded
	if errors.As(err, &limitErr) {
		c.JSON(http.StatusBadRequest, domain.NewPaginationErrorResponse(limitErr.Error()))
		return true
	}

	var uniqueConstraintErr *domain.ErrObjectUniqueConstraint
	if errors.As(err, &uniqueConstraintErr) {
		c.JSON(http.StatusConflict, domain.NewUniqueConstraintErrorResponse(err.Error()))
		return true
	}

	return false
}
