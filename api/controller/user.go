package controller

import (
	"ExerciseManager/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	usecase domain.UserUsecase
}

func NewUserController(usecase domain.UserUsecase) *UserController {
	return &UserController{usecase: usecase}
}

func (uc *UserController) Get(c *gin.Context) {
	uc.usecase.Get(&domain.FilterParams{})
}

func (uc *UserController) List(c *gin.Context) {
	paginationParams, err := getUserPaginationParams(c)
	if handlePaginationLimitExceededError(c, err) {
		return
	}

	params := domain.Params{
		Pagination: paginationParams,
	}

	paginatedResult, err := uc.usecase.List(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.InternalServerError)
		return
	}

	responseUsers := domain.ToResponseUsers(paginatedResult.Data)

	paginatedResponse := domain.PaginatedResponse{
		Count:  int(paginatedResult.Count),
		Limit:  paginationParams.Limit,
		Offset: paginationParams.Offset,
		Data:   responseUsers,
	}

	c.JSON(http.StatusOK, paginatedResponse)
}

func (uc *UserController) Create(c *gin.Context) {
	uc.usecase.Create(&domain.User{})
}

func (uc *UserController) Update(c *gin.Context) {
	uc.usecase.Update(&domain.User{})
}

func (uc *UserController) Delete(c *gin.Context) {
	uc.usecase.Delete(0)
}
