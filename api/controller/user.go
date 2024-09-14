package controller

import (
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/domain"
	"ExerciseManager/internal/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultUserController struct {
	usecase   domain.UserUsecase
	validator validation.UserValidator
	Logger    bootstrap.Logger
}

func NewDefaultUserController(
	usecase domain.UserUsecase,
	validator validation.UserValidator,
	logger bootstrap.Logger,
) *DefaultUserController {
	return &DefaultUserController{
		usecase:   usecase,
		validator: validator,
		Logger:    logger,
	}
}

func (uc *DefaultUserController) Me(c *gin.Context) {
	authUserId := c.GetUint("x-user-id")

	user, err := uc.usecase.GetById(authUserId)

	if err != nil {
		if tryToHandleErr(c, err) {
			return
		}
		c.JSON(http.StatusInternalServerError, domain.InternalServerErrorResponse)
		return
	}

	responseUser := user.ToResponseUser()
	c.JSON(http.StatusOK, responseUser)
}

func (uc *DefaultUserController) Get(c *gin.Context) {
	Id, IsFound := tryToGetIdParamOrBadRequest(c, "id")
	if !IsFound {
		return
	}

	user, err := uc.usecase.GetById(Id)

	if err != nil {
		if tryToHandleErr(c, err) {
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
	if err != nil {
		if tryToHandleErr(c, err) {
			return
		}
		c.JSON(http.StatusInternalServerError, domain.InternalServerErrorResponse)
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
		if tryToHandleErr(c, err) {
			return
		}
		c.JSON(http.StatusInternalServerError, domain.InternalServerErrorResponse)
		return
	}

	responseUser := createdUser.ToResponseUser()

	c.JSON(http.StatusCreated, responseUser)
}

func (uc *DefaultUserController) Update(c *gin.Context) {
	Id, IsFound := tryToGetIdParamOrBadRequest(c, "id")
	if !IsFound {
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

	updatedUser, err := uc.usecase.Update(authUserId, Id, &user)
	if err != nil {
		if tryToHandleErr(c, err) {
			return
		}
		c.JSON(http.StatusInternalServerError, domain.InternalServerErrorResponse)
		return
	}

	responseUser := updatedUser.ToResponseUser()

	c.JSON(http.StatusOK, responseUser)
}

func (uc *DefaultUserController) Delete(c *gin.Context) {
	Id, IsFound := tryToGetIdParamOrBadRequest(c, "id")
	if !IsFound {
		return
	}

	authUserId := c.GetUint("x-user-id")

	if err := uc.usecase.Delete(authUserId, Id); err != nil {
		if tryToHandleErr(c, err) {
			return
		}
		c.JSON(http.StatusInternalServerError, domain.InternalServerErrorResponse)
		return
	}
	c.Status(http.StatusNoContent)
}
