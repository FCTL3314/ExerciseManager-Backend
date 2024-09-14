package router

import (
	"ExerciseManager/api/controller"
	"ExerciseManager/internal/repository"
	"ExerciseManager/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(gin *gin.Engine, db *gorm.DB) {
	v1Router := gin.Group("/api/v1/")

	registerUserRoutes(v1Router, db)
}

func registerUserRoutes(baseRouter *gin.RouterGroup, db *gorm.DB) {
	usersRouter := baseRouter.Group("/users/")

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	userRouter := NewUserRouter(usersRouter, userController)

	userRouter.RegisterAll()
}
