package router

import (
	"ExerciseManager/api/controller"
	"ExerciseManager/api/middleware"
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/auth"
	"ExerciseManager/internal/errormapper"
	"ExerciseManager/internal/permission"
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

	userRepository := repository.NewUserRepository(db)
	accessManager := permission.BuildDefaultAccessManager()
	passwordHasher := auth.NewBcryptPasswordHasher()
	tokenManager := tokenutil.NewJWTTokenManager(
		cfg.JWTAccessSecret,
		cfg.JWTRefreshSecret,
		cfg.JWTAccessExpire,
		cfg.JWTRefreshExpire,
	)
	errorMapper := errormapper.BuildAllErrorsMapperChain()
	userUsecase := usecase.NewUserUsecase(
		userRepository,
		accessManager,
		passwordHasher,
		tokenManager,
		errorMapper,
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
