package router

import (
	"ExerciseManager/api/controller"
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/accesscontrol"
	"ExerciseManager/internal/auth"
	"ExerciseManager/internal/repository"
	"ExerciseManager/internal/usecase"
	"ExerciseManager/internal/validation"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(gin *gin.Engine, db *gorm.DB, cfg *bootstrap.Config) {
	v1Router := gin.Group("/api/v1/")

	registerUserRoutes(v1Router, db, cfg)
}

func registerUserRoutes(baseRouter *gin.RouterGroup, db *gorm.DB, cfg *bootstrap.Config) {
	usersRouter := baseRouter.Group("/users/")

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(
		userRepository,
		accesscontrol.NewUserAccess(),
		auth.NewBcryptPasswordHasher(),
	)
	userController := controller.NewDefaultUserController(
		userUsecase,
		validation.NewDefaultUserValidator(),
	)
	userRouter := NewUserRouter(usersRouter, userController, cfg)

	userRouter.RegisterAll()
}
