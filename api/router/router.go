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

type GetRouter interface {
	RegisterGet()
}

type ListRouter interface {
	RegisterList()
}

type CreateRouter interface {
	RegisterCreate()
}

type UpdateRouter interface {
	RegisterUpdate()
}

type DeleteRouter interface {
	RegisterDelete()
}

type AllRouter interface {
	RegisterAll()
}

type Router interface {
	GetRouter
	ListRouter
	CreateRouter
	UpdateRouter
	DeleteRouter
	AllRouter
}

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
		cfg,
	)
	userController := controller.NewUserController(
		userUsecase,
		validation.NewUserValidator(),
		logger,
	)
	userRouter := NewUserRouter(usersRouter, userController, cfg)

	userRouter.RegisterAll()
}
