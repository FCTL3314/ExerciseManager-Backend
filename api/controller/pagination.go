package controller

import (
	"ExerciseManager/internal/domain"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	MaxUserPaginationLimit = 32
)

func handlePaginationLimitExceededError(c *gin.Context, err error) bool {
	if err != nil {
		var limitErr *domain.PaginationLimitExceededError
		if errors.As(err, &limitErr) {
			c.JSON(http.StatusBadRequest, domain.NewPaginationErrorResponse(limitErr.Error()))
			return true
		}

		c.JSON(http.StatusInternalServerError, domain.InternalServerErrorResponse)
		return true
	}
	return false
}

func getPaginationParams(c *gin.Context, maxLimit int) (domain.PaginationParams, error) {
	limitStr := c.DefaultQuery("limit", strconv.Itoa(maxLimit))
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	} else if limit > maxLimit {
		return domain.PaginationParams{}, &domain.PaginationLimitExceededError{MaxLimit: maxLimit}
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	return domain.PaginationParams{
		Limit:  limit,
		Offset: offset,
	}, nil
}

func getUserPaginationParams(c *gin.Context) (domain.PaginationParams, error) {
	return getPaginationParams(c, MaxUserPaginationLimit)
}
