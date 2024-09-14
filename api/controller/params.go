package controller

import (
	"ExerciseManager/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func tryToGetIdParamOrBadRequest(c *gin.Context, param string) (Id uint, IsFound bool) {
	id, err := strconv.ParseUint(c.Param(param), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.InvalidURLParamErrorResponse)
		return 0, false
	}
	return uint(id), true
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
