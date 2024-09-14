package router

import (
	"ExerciseManager/api/controller"
	"ExerciseManager/api/middleware"
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/accesscontrol"
	"ExerciseManager/internal/auth"
	"ExerciseManager/internal/repository"
	"ExerciseManager/internal/usecase"
	"ExerciseManager/internal/validation"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(
	gin *gin.Engine,
	db *gorm.DB,
	cfg *bootstrap.Config,
	loggerGroup *bootstrap.LoggerGroup,
) {
	v1Router := gin.Group("/api/v1/")
	registerUserRoutes(v1Router, db, cfg, *loggerGroup.User)
}

func registerUserRoutes(
	baseRouter *gin.RouterGroup,
	db *gorm.DB,
	cfg *bootstrap.Config,
	logger bootstrap.Logger,
) {
	usersRouter := baseRouter.Group("/users/")
	usersRouter.Use(middleware.ErrorLoggerMiddleware(logger))

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(
		userRepository,
		accesscontrol.NewUserAccess(),
		auth.NewBcryptPasswordHasher(),
	)
	userController := controller.NewDefaultUserController(
		userUsecase,
		validation.NewDefaultUserValidator(),
		logger,
	)
	userRouter := NewUserRouter(usersRouter, userController, cfg)

	userRouter.RegisterAll()
}
