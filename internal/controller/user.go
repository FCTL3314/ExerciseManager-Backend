package controller

import (
	"ExerciseManager/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase domain.UserUsecase
}

func NewUserController(userUsecase domain.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

func (uc *UserController) Get(c *gin.Context) (*domain.User, error) {
	return uc.userUsecase.Get(&domain.FilterParams{})
}

func (uc *UserController) List(c *gin.Context) ([]*domain.User, error) {
	return uc.userUsecase.List(&domain.Params{})
}

func (uc *UserController) Create(c *gin.Context) (*domain.User, error) {
	return uc.userUsecase.Create(&domain.User{})
}

func (uc *UserController) Delete(c *gin.Context) error {
	return uc.userUsecase.Delete(0)
}
