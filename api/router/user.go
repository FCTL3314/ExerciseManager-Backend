package router

import (
	"ExerciseManager/api/controller"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	router         *gin.RouterGroup
	userController *controller.UserController
}

func NewUserRouter(
	router *gin.RouterGroup,
	userController *controller.UserController,
) *UserRouter {
	return &UserRouter{router, userController}
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
	ur.router.PATCH("/:id", ur.userController.Update)
}

func (ur *UserRouter) RegisterDelete() {
	ur.router.DELETE("/:id", ur.userController.Delete)
}
