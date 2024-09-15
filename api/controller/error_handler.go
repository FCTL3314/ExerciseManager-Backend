package controller

import (
	"ExerciseManager/internal/domain"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorHandlerFunc = func(c *gin.Context, err error) bool

func handleObjectNotFound(c *gin.Context, err error) bool {
	if errors.Is(err, domain.ErrObjectNotFound) {
		c.JSON(http.StatusNotFound, domain.NotFoundResponse)
		return true
	}
	return false
}

func handleAccessDenied(c *gin.Context, err error) bool {
	if errors.Is(err, domain.ErrAccessDenied) {
		c.JSON(http.StatusForbidden, domain.ForbiddenResponse)
		return true
	}
	return false
}

func handleInvalidParam(c *gin.Context, err error) bool {
	var errInvalidURLParam *domain.ErrInvalidURLParam
	if errors.As(err, &errInvalidURLParam) {
		c.JSON(http.StatusBadRequest, domain.NewInvalidURLParamResponse(errInvalidURLParam.Error()))
		return true
	}
	return false
}

func handlePaginationLimitExceeded(c *gin.Context, err error) bool {
	var errPaginationLimit *domain.ErrPaginationLimitExceeded
	if errors.As(err, &errPaginationLimit) {
		c.JSON(http.StatusBadRequest, domain.NewPaginationErrorResponse(errPaginationLimit.Error()))
		return true
	}
	return false
}

func handleAuthInvalidCredentials(c *gin.Context, err error) bool {
	if errors.Is(err, domain.ErrInvalidAuthCredentials) {
		c.JSON(http.StatusUnauthorized, domain.InvalidAuthCredentialsResponse)
		return true
	}
	return false
}

func handleUniqueConstraint(c *gin.Context, err error) bool {
	var errObjectUniqueConstraint *domain.ErrObjectUniqueConstraint
	if errors.As(err, &errObjectUniqueConstraint) {
		c.JSON(http.StatusUnauthorized, domain.NewUniqueConstraintErrorResponse(err.Error()))
		return true
	}
	return false
}

type ErrorHandler struct {
	handlers []ErrorHandlerFunc
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (eh *ErrorHandler) RegisterHandler(handler ErrorHandlerFunc) {
	eh.handlers = append(eh.handlers, handler)
}

func (eh *ErrorHandler) Handle(c *gin.Context, err error) {
	for _, handler := range eh.handlers {
		if handler(c, err) {
			return
		}
	}
	c.JSON(http.StatusInternalServerError, domain.InternalServerErrorResponse)
}

func DefaultErrorHandler() *ErrorHandler {
	eh := NewErrorHandler()
	eh.RegisterHandler(handleObjectNotFound)
	eh.RegisterHandler(handleAccessDenied)
	eh.RegisterHandler(handleInvalidParam)
	eh.RegisterHandler(handlePaginationLimitExceeded)
	return eh
}

func UserErrorHandler() *ErrorHandler {
	eh := DefaultErrorHandler()
	eh.RegisterHandler(handleAuthInvalidCredentials)
	eh.RegisterHandler(handleUniqueConstraint)
	return eh
}
