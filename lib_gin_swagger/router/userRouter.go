package router

import (
	"github.com/gin-gonic/gin"
	"lib_gin_swagger/controller"
)

func userRouters(r *gin.RouterGroup) {
	group := r.Group("/user")
	{
		group.POST("/login", controller.UserLogin)
		group.GET("/logout", controller.UserLogout)
	}
}
