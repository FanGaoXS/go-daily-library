package routers

import "daily_library/lib_goth_gin/controller"

func IndexRouters() {
	group := r.Group("")
	{
		group.GET("", controller.Welcome)
	}
}
