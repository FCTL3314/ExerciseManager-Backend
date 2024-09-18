package controller

import (
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultUserController struct {
	usecase      domain.UserUsecase
	errorHandler *ErrorHandler
	Logger       bootstrap.Logger
}

func NewUserController(
	usecase domain.UserUsecase,
	errorHandler *ErrorHandler,
	logger bootstrap.Logger,
) *DefaultUserController {
	return &DefaultUserController{
		usecase:      usecase,
		errorHandler: errorHandler,
		Logger:       logger,
	}
}

func (uc *DefaultUserController) Me(c *gin.Context) {
	authUserId := c.GetInt64(string(UserIDContextKey))

	user, err := uc.usecase.GetById(authUserId)

	if err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}

	responseUser := user.ToResponseUser()
	c.JSON(http.StatusOK, responseUser)
}

func (uc *DefaultUserController) Get(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}

	user, err := uc.usecase.GetById(id)

	if err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}

	responseUser := user.ToResponseUser()
	c.JSON(http.StatusOK, responseUser)
}

func (uc *DefaultUserController) List(c *gin.Context) {
	paginationParams, err := getUserPaginationParams(c)
	if err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}

	params := domain.Params{
		Pagination: paginationParams,
	}

	paginatedResult, err := uc.usecase.List(&params)
	if err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}

	responseUsers := domain.ToResponseUsers(paginatedResult.Results)

	paginatedResponse := domain.PaginatedResponse{
		Count:   paginatedResult.Count,
		Limit:   paginationParams.Limit,
		Offset:  paginationParams.Offset,
		Results: responseUsers,
	}

	c.JSON(http.StatusOK, paginatedResponse)
}

func (uc *DefaultUserController) Create(c *gin.Context) {
	var user domain.CreateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	createdUser, err := uc.usecase.Create(&user)
	if err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}

	responseUser := createdUser.ToResponseUser()

	c.JSON(http.StatusCreated, responseUser)
}

func (uc *DefaultUserController) Login(c *gin.Context) {
	var user domain.LoginUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	loginResponse, err := uc.usecase.Login(&user)
	if err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}

	c.JSON(http.StatusOK, loginResponse)
}

func (uc *DefaultUserController) RefreshTokens(c *gin.Context) {
	var refreshTokenRequest domain.RefreshTokenRequest
	if err := c.ShouldBindJSON(&refreshTokenRequest); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	refreshTokenResponse, err := uc.usecase.RefreshTokens(&refreshTokenRequest)
	if err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}

func (uc *DefaultUserController) Update(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}

	var user domain.UpdateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	authUserId := c.GetInt64(string(UserIDContextKey))

	updatedUser, err := uc.usecase.Update(authUserId, id, &user)
	if err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}

	responseUser := updatedUser.ToResponseUser()

	c.JSON(http.StatusOK, responseUser)
}

func (uc *DefaultUserController) Delete(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}

	authUserId := c.GetInt64(string(UserIDContextKey))

	if err := uc.usecase.Delete(authUserId, id); err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
