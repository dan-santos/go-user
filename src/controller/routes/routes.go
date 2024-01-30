package routes

import (
	"github.com/dan-santos/go-user/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	controller controller.UserControllerInterface,
) {
	r.GET("/users/:userId", controller.FindUserById)
	r.GET("/users/email/:userEmail", controller.FindUserByEmail)
	r.POST("/users", controller.Create)
	r.PUT("/users/:userId", controller.Update)
	r.DELETE("/users/:userId", controller.Delete)
}