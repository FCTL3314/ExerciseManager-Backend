package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(gin *gin.Engine, db *gorm.DB) {
	baseRouter := gin.Group("/api/v1/")

	registerUserRoutes(baseRouter, db)
}

func registerUserRoutes(baseRouter *gin.RouterGroup, db *gorm.DB) {
	usersRouter := baseRouter.Group("/users/")

	userRouter := NewUserRouter(usersRouter, db)
	userRouter.RegisterAll()
}
