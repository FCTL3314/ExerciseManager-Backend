package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(gin *gin.Engine, db *gorm.DB) {
	baseRouter := gin.Group("/api/")

	registerUser(baseRouter, db)
}

func registerUser(baseRouter *gin.RouterGroup, db *gorm.DB) {
	usersRouter := baseRouter.Group("/users/")

	userRouter := NewUserRouter(usersRouter, db)
	userRouter.RegisterGet()
	userRouter.RegisterList()
	userRouter.RegisterCreate()
	userRouter.RegisterUpdate()
	userRouter.RegisterDelete()

}
