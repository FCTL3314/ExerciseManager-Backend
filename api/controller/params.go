package controller

import (
	"ExerciseManager/internal/domain"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ContextKey string

const (
	UserIDContextKey ContextKey = "userID"
)

func getParamAsInt64(c *gin.Context, key string) (int64, error) {
	id, err := strconv.ParseInt(c.Param(key), 10, 64)
	if err != nil {
		return 0, domain.ErrInvalidParam
	}
	return id, nil
}

func getPaginationParams(c *gin.Context, maxLimit int) (domain.PaginationParams, error) {
	limitStr := c.DefaultQuery("limit", strconv.Itoa(maxLimit))
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	} else if limit > maxLimit {
		return domain.PaginationParams{}, &domain.ErrPaginationLimitExceeded{MaxLimit: maxLimit}
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
	return getPaginationParams(c, domain.MaxUserPaginationLimit)
}
