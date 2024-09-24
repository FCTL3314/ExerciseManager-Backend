package controller

import (
	"ExerciseManager/internal/collections"
	"ExerciseManager/internal/domain"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ContextKey string

const (
	UserIDContextKey ContextKey = "userID"
)

var (
	FilterParamsToExclude = []string{"limit", "offset"}
)

func getParamAsInt64(c *gin.Context, key string) (int64, error) {
	id, err := strconv.ParseInt(c.Param(key), 10, 64)
	if err != nil {
		return 0, &domain.ErrInvalidURLParam{Param: key}
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

func getFilterParams(c *gin.Context) (domain.FilterParams, error) {
	queryParams := c.Request.URL.Query()
	filter := domain.FilterParams{
		Query: "",
		Args:  []interface{}{},
	}

	for key, values := range queryParams {
		if collections.Contains(FilterParamsToExclude, key) {
			continue
		}

		for _, value := range values {
			query, ok := filter.Query.(string)
			if !ok {
				query = ""
			}

			if query != "" {
				query += " AND "
			}
			query += key + " = ?"
			filter.Query = query
			filter.Args = append(filter.Args, value)
		}
	}

	return filter, nil
}

func getParams(c *gin.Context, paginationMaxLimit int) (domain.Params, error) {
	paginationParams, err := getPaginationParams(c, paginationMaxLimit)
	if err != nil {
		return domain.Params{}, err
	}

	filterParams, err := getFilterParams(c)
	if err != nil {
		return domain.Params{}, err
	}

	return domain.Params{
		Pagination: paginationParams,
		Filter:     filterParams,
	}, nil

}
