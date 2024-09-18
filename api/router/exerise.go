package router

import (
	"ExerciseManager/api/controller"
	"ExerciseManager/api/middleware"
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

type ExerciseRouter struct {
	router             *gin.RouterGroup
	tokenManager       tokenutil.JWTTokenManager
	exerciseController controller.Controller
	cfg                *bootstrap.Config
}

func NewExerciseRouter(
	router *gin.RouterGroup,
	tokenManager *tokenutil.DefaultJWTTokenManager,
	exerciseController *controller.DefaultExerciseController,
	cfg *bootstrap.Config,
) *ExerciseRouter {
	return &ExerciseRouter{router, tokenManager, exerciseController, cfg}
}

func (wr *ExerciseRouter) RegisterAll() {
	wr.RegisterGet()
	wr.RegisterList()
	wr.RegisterCreate()
	wr.RegisterUpdate()
	wr.RegisterDelete()
}

func (wr *ExerciseRouter) RegisterGet() {
	wr.router.GET("/:id", wr.exerciseController.Get)
}

func (wr *ExerciseRouter) RegisterList() {
	wr.router.GET("", wr.exerciseController.List)
}

func (wr *ExerciseRouter) RegisterCreate() {
	wr.router.POST("", middleware.JwtAuthMiddleware(wr.tokenManager), wr.exerciseController.Create)
}

func (wr *ExerciseRouter) RegisterUpdate() {
	wr.router.PATCH("/:id", middleware.JwtAuthMiddleware(wr.tokenManager), wr.exerciseController.Update)
}

func (wr *ExerciseRouter) RegisterDelete() {
	wr.router.DELETE("/:id", middleware.JwtAuthMiddleware(wr.tokenManager), wr.exerciseController.Delete)
}
