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
	registerWorkoutRoutes(v1Router, db, cfg, *loggerGroup.Workout)
	registerExerciseRoutes(v1Router, db, cfg, *loggerGroup.Exercise)
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
	accessManager := permission.BuildUserAccessManager()
	passwordHasher := auth.NewBcryptPasswordHasher()
	tokenManager := tokenutil.NewJWTTokenManager(
		cfg.JWT.AccessSecret,
		cfg.JWT.RefreshSecret,
		cfg.JWT.AccessExpire,
		cfg.JWT.RefreshExpire,
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
		cfg,
	)

	userRouter := NewUserRouter(usersRouter, tokenManager, userController, cfg)
	userRouter.RegisterAll()
}

func registerWorkoutRoutes(
	baseRouter *gin.RouterGroup,
	db *gorm.DB,
	cfg *bootstrap.Config,
	logger bootstrap.Logger,
) {
	workoutsRouter := baseRouter.Group("/workouts/")
	workoutsRouter.Use(middleware.ErrorLoggerMiddleware(logger))

	workoutRepository := repository.NewWorkoutRepository(db)
	exerciseRepository := repository.NewExerciseRepository(db)
	workoutExerciseRepository := repository.NewWorkoutExerciseRepository(db)
	accessManager := permission.BuildWorkoutAccessManager()
	tokenManager := tokenutil.NewJWTTokenManager(
		cfg.JWT.AccessSecret,
		cfg.JWT.RefreshSecret,
		cfg.JWT.AccessExpire,
		cfg.JWT.RefreshExpire,
	)
	errorMapper := errormapper.BuildAllErrorsMapperChain()
	workoutUsecase := usecase.NewWorkoutUsecase(
		workoutRepository,
		exerciseRepository,
		workoutExerciseRepository,
		accessManager,
		errorMapper,
		cfg,
	)

	errorHandler := controller.DefaultErrorHandler()
	workoutController := controller.NewWorkoutController(
		workoutUsecase,
		errorHandler,
		logger,
		cfg,
	)

	workoutRouter := NewWorkoutRouter(workoutsRouter, tokenManager, workoutController, cfg)
	workoutRouter.RegisterAll()
}

func registerExerciseRoutes(
	baseRouter *gin.RouterGroup,
	db *gorm.DB,
	cfg *bootstrap.Config,
	logger bootstrap.Logger,
) {
	exercisesRouter := baseRouter.Group("/exercises/")
	exercisesRouter.Use(middleware.ErrorLoggerMiddleware(logger))

	exerciseRepository := repository.NewExerciseRepository(db)
	accessManager := permission.BuildExerciseAccessManager()
	tokenManager := tokenutil.NewJWTTokenManager(
		cfg.JWT.AccessSecret,
		cfg.JWT.RefreshSecret,
		cfg.JWT.AccessExpire,
		cfg.JWT.RefreshExpire,
	)
	errorMapper := errormapper.BuildAllErrorsMapperChain()
	exerciseUsecase := usecase.NewExerciseUsecase(
		exerciseRepository,
		accessManager,
		errorMapper,
	)

	errorHandler := controller.DefaultErrorHandler()
	exerciseController := controller.NewExerciseController(
		exerciseUsecase,
		errorHandler,
		logger,
		cfg,
	)

	exerciseRouter := NewExerciseRouter(exercisesRouter, tokenManager, exerciseController, cfg)
	exerciseRouter.RegisterAll()
}
