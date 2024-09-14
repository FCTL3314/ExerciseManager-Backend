package router

import (
	"ExerciseManager/api/controller"
	"ExerciseManager/api/middleware"
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/accesscontrol"
	"ExerciseManager/internal/repository"
	"ExerciseManager/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(gin *gin.Engine, db *gorm.DB, cfg *bootstrap.Config) {
	v1Router := gin.Group("/api/v1/")

	registerUserRoutes(v1Router, db, cfg)
}

func registerUserRoutes(baseRouter *gin.RouterGroup, db *gorm.DB, cfg *bootstrap.Config) {
	usersRouter := baseRouter.Group("/users/")
	usersRouter.Use(middleware.JwtAuthMiddleware(cfg.JWTSecret))

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, &accesscontrol.User{})
	userController := controller.NewUserController(userUsecase)
	userRouter := NewUserRouter(usersRouter, userController)

	userRouter.RegisterAll()
}
