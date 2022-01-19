package routers

import "github.com/gin-gonic/gin"

const (
	Port = ":9090"
)

var (
	r *gin.Engine
)

func Init() {
	r = gin.Default()
	routers()
	r.Run(Port)
}

func routers() {
	IndexRouters()
	AuthRouters()
}
