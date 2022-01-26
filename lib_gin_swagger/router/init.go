package router

import (
	"github.com/gin-gonic/gin"
)

const (
	Port = "9090"
)

var (
	r *gin.Engine
)

func InitGin() {
	r = gin.Default()
	registerRouters(v1())
	initSwagger()
	r.Run(":" + Port)
}

func registerRouters(r *gin.RouterGroup) {
	pingRouters()
	userRouters(r)
}

func GetRouterEngine() *gin.Engine {
	return r
}

func v1() *gin.RouterGroup {
	return r.Group("/api/v1")
}
