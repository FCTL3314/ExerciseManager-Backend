package controller

import (
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	usecase      domain.UserUsecase
	errorHandler *ErrorHandler
	Logger       bootstrap.Logger
}

func NewUserController(
	usecase domain.UserUsecase,
	errorHandler *ErrorHandler,
	logger bootstrap.Logger,
) *UserController {
	return &UserController{
		usecase:      usecase,
		errorHandler: errorHandler,
		Logger:       logger,
	}
}

func (uc *UserController) Me(c *gin.Context) {
	authUserId := c.GetUint(string(UserIDContextKey))

	user, err := uc.usecase.GetById(authUserId)

	if err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}

	responseUser := user.ToResponseUser()
	c.JSON(http.StatusOK, responseUser)
}

func (uc *UserController) Get(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}

	user, err := uc.usecase.GetById(uint(id))

	if err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}

	responseUser := user.ToResponseUser()
	c.JSON(http.StatusOK, responseUser)
}

func (uc *UserController) List(c *gin.Context) {
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
		Count:   int(paginatedResult.Count),
		Limit:   paginationParams.Limit,
		Offset:  paginationParams.Offset,
		Results: responseUsers,
	}

	c.JSON(http.StatusOK, paginatedResponse)
}

func (uc *UserController) Create(c *gin.Context) {
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

func (uc *UserController) Login(c *gin.Context) {
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

func (uc *UserController) RefreshTokens(c *gin.Context) {
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

func (uc *UserController) Update(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}

	authUserId := c.GetUint(string(UserIDContextKey))

	var user domain.UpdateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	updatedUser, err := uc.usecase.Update(authUserId, uint(id), &user)
	if err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}

	responseUser := updatedUser.ToResponseUser()

	c.JSON(http.StatusOK, responseUser)
}

func (uc *UserController) Delete(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}

	authUserId := c.GetUint(string(UserIDContextKey))

	if err := uc.usecase.Delete(authUserId, uint(id)); err != nil {
		uc.errorHandler.Handle(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
