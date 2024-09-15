package router

import (
	"ExerciseManager/api/controller"
	"ExerciseManager/api/middleware"
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	router         *gin.RouterGroup
	tokenManager   *tokenutil.TokenManager
	userController controller.IUserController
	cfg            *bootstrap.Config
}

func NewUserRouter(
	router *gin.RouterGroup,
	tokenManager *tokenutil.TokenManager,
	userController *controller.UserController,
	cfg *bootstrap.Config,
) *UserRouter {
	return &UserRouter{router, tokenManager, userController, cfg}
}

func (ur *UserRouter) RegisterAll() {
	ur.RegisterMe()
	ur.RegisterGet()
	ur.RegisterList()
	ur.RegisterCreate()
	ur.RegisterLogin()
	ur.RegisterRefreshTokens()
	ur.RegisterUpdate()
	ur.RegisterDelete()
}

func (ur *UserRouter) RegisterMe() {
	ur.router.GET("/me", middleware.JwtAuthMiddleware(ur.tokenManager), ur.userController.Me)
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

func (ur *UserRouter) RegisterLogin() {
	ur.router.POST("/login", ur.userController.Login)
}

func (ur *UserRouter) RegisterRefreshTokens() {
	ur.router.POST("/refresh", ur.userController.RefreshTokens)
}

func (ur *UserRouter) RegisterUpdate() {
	ur.router.PATCH("/:id", middleware.JwtAuthMiddleware(ur.tokenManager), ur.userController.Update)
}

func (ur *UserRouter) RegisterDelete() {
	ur.router.DELETE("/:id", middleware.JwtAuthMiddleware(ur.tokenManager), ur.userController.Delete)
}
