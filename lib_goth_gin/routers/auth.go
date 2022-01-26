package routers

import "lib_goth_gin/controller"

func AuthRouters() {
	group := r.Group("/auth")
	{
		group.GET("", controller.BeginAuth)
		group.GET("/callback/:provider", controller.AuthCallBack)
		group.GET("/logout/:provider", controller.Logout)
	}
}
