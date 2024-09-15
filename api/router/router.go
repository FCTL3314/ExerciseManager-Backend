package router

import (
	"ExerciseManager/api/controller"
	"ExerciseManager/api/middleware"
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/accesscontrol"
	"ExerciseManager/internal/auth"
	"ExerciseManager/internal/repository"
	"ExerciseManager/internal/tokenutil"
	"ExerciseManager/internal/usecase"
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

	tokenManager := tokenutil.NewJWTTokenManager(
		cfg.JWTAccessSecret, cfg.JWTRefreshSecret, cfg.JWTAccessExpire, cfg.JWTRefreshExpire,
	)

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(
		userRepository,
		accesscontrol.NewUserAccess(),
		auth.NewBcryptPasswordHasher(),
		tokenManager,
	)
	errorHandler := controller.UserErrorHandler()
	userController := controller.NewUserController(
		userUsecase,
		errorHandler,
		logger,
	)
	userRouter := NewUserRouter(usersRouter, tokenManager, userController, cfg)

	userRouter.RegisterAll()
}
