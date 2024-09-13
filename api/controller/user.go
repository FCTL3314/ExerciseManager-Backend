package controller

import (
	"ExerciseManager/internal/domain"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
	usecase domain.UserUsecase
}

func NewUserController(usecase domain.UserUsecase) *UserController {
	return &UserController{usecase: usecase}
}

func (uc *UserController) Get(c *gin.Context) {
	user, err := uc.usecase.Get(
		&domain.FilterParams{
			Query: "id = ?",
			Args:  []interface{}{c.Param("id")},
		},
	)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, domain.NotFoundResponse)
			return
		}
		c.JSON(http.StatusInternalServerError, domain.InternalServerErrorResponse)
		return
	}

	c.JSON(http.StatusOK, user)
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
		c.JSON(http.StatusInternalServerError, domain.InternalServerErrorResponse)
		return
	}

	responseUsers := domain.ToResponseUsers(paginatedResult.Results)

	paginatedResponse := domain.PaginatedResponse{
		Count:   int(paginatedResult.Count),
		Limit:   paginationParams.Limit,
		Offset:  paginationParams.Offset,
		Results: responseUsers,
	}

	c.JSON(http.StatusOK, paginatedResponse)
}

func (uc *UserController) Create(c *gin.Context) {
	var user domain.CreateUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	createdUser, err := uc.usecase.Create(&user)
	if err != nil {

		var uniqueConstraintErr *domain.ObjectUniqueConstraintError
		if errors.As(err, &uniqueConstraintErr) {
			c.JSON(http.StatusConflict, domain.NewUniqueConstraintErrorResponse(err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, domain.InternalServerErrorResponse)
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (uc *UserController) Update(c *gin.Context) {
	uc.usecase.Update(&domain.User{})
}

func (uc *UserController) Delete(c *gin.Context) {
	uc.usecase.Delete(0)
}