package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func pingRouters() {
	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
