package controller

import (
	"ExerciseManager/bootstrap"
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

	var limitErr *domain.ErrPaginationLimitExceeded
	if errors.As(err, &limitErr) {
		c.JSON(http.StatusBadRequest, domain.NewPaginationErrorResponse(limitErr.Error()))
		return true
	}

	var uniqueConstraintErr *domain.ErrObjectUniqueConstraint
	if errors.As(err, &uniqueConstraintErr) {
		c.JSON(http.StatusConflict, domain.NewUniqueConstraintErrorResponse(err.Error()))
		return
	}

	return false
}

func tryToHandleErrOrLog(c *gin.Context, err error, logger bootstrap.Logger) (IsHandled bool) {
	if !tryToHandleErr(c, err) {
		logger.Error(err.Error())
		return false
	}
	return true
}
