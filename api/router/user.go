package router

import (
	"ExerciseManager/api/controller"
	"ExerciseManager/internal/repository"
	"ExerciseManager/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRouter struct {
	router         *gin.RouterGroup
	db             *gorm.DB
	userController *controller.UserController
}

func NewUserRouter(router *gin.RouterGroup, db *gorm.DB) *UserRouter {
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	return &UserRouter{router, db, userController}
}

func (ur *UserRouter) RegisterAll() {
	ur.RegisterGet()
	ur.RegisterList()
	ur.RegisterCreate()
	ur.RegisterUpdate()
	ur.RegisterDelete()
}

func (ur *UserRouter) RegisterGet() {
	ur.router.GET("/:id", ur.userController.Get)
}

func (ur *UserRouter) RegisterList() {
	ur.router.GET("", ur.userController.List)
}

func (ur *UserRouter) RegisterCreate() {
	ur.router.POST("", ur.userController.Create)
}

func (ur *UserRouter) RegisterUpdate() {
	ur.router.PATCH("", ur.userController.Update)
}

func (ur *UserRouter) RegisterDelete() {
	ur.router.DELETE("", ur.userController.Delete)
}
