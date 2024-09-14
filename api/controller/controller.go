package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	Get(c *gin.Context)
	List(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type IUserController interface {
	Me(c *gin.Context)
	Login(c *gin.Context)
	Controller
}
