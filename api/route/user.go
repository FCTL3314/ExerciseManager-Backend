package route

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

func (ur *UserRouter) RegisterGet() {

}

func (ur *UserRouter) RegisterList() {
	ur.router.GET("", ur.userController.List)
}

func (ur *UserRouter) RegisterCreate() {

}

func (ur *UserRouter) RegisterUpdate() {

}

func (ur *UserRouter) RegisterDelete() {

}
