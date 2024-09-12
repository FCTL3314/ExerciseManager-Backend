package controller

import (
	"ExerciseManager/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userUsecase domain.UserUsecase
}

func NewUserController(userUsecase domain.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

func (uc *UserController) Get(c *gin.Context) {
	uc.userUsecase.Get(&domain.FilterParams{})
}

func (uc *UserController) List(c *gin.Context) {
	users, err := uc.userUsecase.List(&domain.Params{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uc *UserController) Create(c *gin.Context) {
	uc.userUsecase.Create(&domain.User{})
}

func (uc *UserController) Delete(c *gin.Context) {
	uc.userUsecase.Delete(0)
}
