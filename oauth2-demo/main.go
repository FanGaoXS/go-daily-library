package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"oauth-demo/service_provider/gitee"
	"oauth-demo/service_provider/github"
	"oauth-demo/service_provider/gitlab"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	g1 := r.Group("/github")
	{
		g1.GET("/authorize", github.Authorize)
		g1.GET("/oauth/callback", github.Callback)
		g1.GET("/userinfo", github.Userinfo)
	}

	g2 := r.Group("/gitee")
	{
		g2.GET("/authorize", gitee.Authorize)
		g2.GET("/oauth/callback", gitee.Callback)
		g2.GET("/userinfo", gitee.Userinfo)
	}

	g3 := r.Group("/gitlab")
	{
		g3.GET("/authorize", gitlab.Authorize)
		g3.GET("/oauth/callback", gitlab.Callback)
		g3.GET("/userinfo", gitlab.Userinfo)
	}

	r.Run(":8090")
}
