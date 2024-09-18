package router

import (
	"ExerciseManager/api/controller"
	"ExerciseManager/api/middleware"
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

type WorkoutRouter struct {
	router            *gin.RouterGroup
	tokenManager      tokenutil.JWTTokenManager
	workoutController controller.Controller
	cfg               *bootstrap.Config
}

func NewWorkoutRouter(
	router *gin.RouterGroup,
	tokenManager *tokenutil.DefaultJWTTokenManager,
	workoutController *controller.DefaultWorkoutController,
	cfg *bootstrap.Config,
) *WorkoutRouter {
	return &WorkoutRouter{router, tokenManager, workoutController, cfg}
}

func (wr *WorkoutRouter) RegisterAll() {
	wr.RegisterGet()
	wr.RegisterList()
	wr.RegisterCreate()
	wr.RegisterUpdate()
	wr.RegisterDelete()
}

func (wr *WorkoutRouter) RegisterGet() {
	wr.router.GET("/:id", wr.workoutController.Get)
}

func (wr *WorkoutRouter) RegisterList() {
	wr.router.GET("", wr.workoutController.List)
}

func (wr *WorkoutRouter) RegisterCreate() {
	wr.router.POST("", middleware.JwtAuthMiddleware(wr.tokenManager), wr.workoutController.Create)
}

func (wr *WorkoutRouter) RegisterUpdate() {
	wr.router.PATCH("/:id", middleware.JwtAuthMiddleware(wr.tokenManager), wr.workoutController.Update)
}

func (wr *WorkoutRouter) RegisterDelete() {
	wr.router.DELETE("/:id", middleware.JwtAuthMiddleware(wr.tokenManager), wr.workoutController.Delete)
}