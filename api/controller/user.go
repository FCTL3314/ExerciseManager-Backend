package controller

import (
	"ExerciseManager/internal/domain"
	"ExerciseManager/internal/validation"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DefaultUserController struct {
	usecase   domain.UserUsecase
	validator validation.UserValidator
}

func NewDefaultUserController(
	usecase domain.UserUsecase,
	validator validation.UserValidator,
) *DefaultUserController {
	return &DefaultUserController{usecase: usecase, validator: validator}
}

func (uc *DefaultUserController) Me(c *gin.Context) {
	authUserId := c.GetUint("x-user-id")

	user, err := uc.usecase.GetById(authUserId)

	if err != nil {
		if errors.Is(err, domain.ErrObjectNotFound) {
			c.JSON(http.StatusNotFound, domain.NotFoundResponse)
			return
		}
		c.JSON(http.StatusInternalServerError, domain.InternalServerErrorResponse)
		return
	}

	responseUser := user.ToResponseUser()

	c.JSON(http.StatusOK, responseUser)
}

func (uc *DefaultUserController) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.InvalidURLParamErrorResponse)
		return
	}

	user, err := uc.usecase.GetById(uint(id))

	if err != nil {
		if errors.Is(err, domain.ErrObjectNotFound) {
			c.JSON(http.StatusNotFound, domain.NotFoundResponse)
			return
		}
		c.JSON(http.StatusInternalServerError, domain.InternalServerErrorResponse)
		return
	}

	responseUser := user.ToResponseUser()

	c.JSON(http.StatusOK, responseUser)
}

func (uc *DefaultUserController) List(c *gin.Context) {
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

func (uc *DefaultUserController) Create(c *gin.Context) {
	var user domain.CreateUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	if err := uc.validator.ValidateCreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	createdUser, err := uc.usecase.Create(&user)
	if err != nil {

		var uniqueConstraintErr *domain.ErrObjectUniqueConstraint
		if errors.As(err, &uniqueConstraintErr) {
			c.JSON(http.StatusConflict, domain.NewUniqueConstraintErrorResponse(err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, domain.InternalServerErrorResponse)
		return
	}

	responseUser := createdUser.ToResponseUser()

	c.JSON(http.StatusCreated, responseUser)
}

func (uc *DefaultUserController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.InvalidURLParamErrorResponse)
		return
	}

	authUserId := c.GetUint("x-user-id")

	var user domain.UpdateUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	if err := uc.validator.ValidateUpdateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	updatedUser, err := uc.usecase.Update(uint(authUserId), uint(id), &user)
	if err != nil {
		if errors.Is(err, domain.ErrObjectNotFound) {
			c.JSON(http.StatusNotFound, domain.NotFoundResponse)
			return
		} else if errors.Is(err, domain.ErrAccessDenied) {
			c.JSON(http.StatusForbidden, domain.ForbiddenResponse)
			return
		}
		c.JSON(http.StatusInternalServerError, domain.InternalServerErrorResponse)
		return
	}

	responseUser := updatedUser.ToResponseUser()

	c.JSON(http.StatusOK, responseUser)
}

func (uc *DefaultUserController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.InvalidURLParamErrorResponse)
		return
	}

	authUserIdString := c.GetString("x-user-id")
	authUserId, err := strconv.ParseUint(authUserIdString, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.InternalServerErrorResponse)
	}

	if err := uc.usecase.Delete(uint(authUserId), uint(id)); err != nil {
		if errors.Is(err, domain.ErrObjectNotFound) {
			c.JSON(http.StatusNotFound, domain.NotFoundResponse)
			return
		} else if errors.Is(err, domain.ErrAccessDenied) {
			c.JSON(http.StatusForbidden, domain.ForbiddenResponse)
			return
		}
		c.JSON(http.StatusInternalServerError, domain.InternalServerErrorResponse)
	}
	c.Status(http.StatusNoContent)
}
