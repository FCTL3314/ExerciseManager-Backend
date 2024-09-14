package router

import (
	"ExerciseManager/api/controller"
	"ExerciseManager/api/middleware"
	"ExerciseManager/bootstrap"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	router         *gin.RouterGroup
	userController controller.UserController
	cfg            *bootstrap.Config
}

func NewUserRouter(
	router *gin.RouterGroup,
	userController *controller.DefaultUserController,
	cfg *bootstrap.Config,
) *UserRouter {
	return &UserRouter{router, userController, cfg}
}

func (ur *UserRouter) RegisterAll() {
	ur.RegisterMe()
	ur.RegisterGet()
	ur.RegisterList()
	ur.RegisterCreate()
	ur.RegisterUpdate()
	ur.RegisterDelete()
}

func (ur *UserRouter) RegisterMe() {
	ur.router.GET("/me", middleware.JwtAuthMiddleware(ur.cfg.JWTSecret), ur.userController.Me)
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
	ur.router.PATCH("/:id", middleware.JwtAuthMiddleware(ur.cfg.JWTSecret), ur.userController.Update)
}

func (ur *UserRouter) RegisterDelete() {
	ur.router.DELETE("/:id", middleware.JwtAuthMiddleware(ur.cfg.JWTSecret), ur.userController.Delete)
}
