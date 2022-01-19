package main

import (
	"daily_library/lib_goth_gin/providers"
	"daily_library/lib_goth_gin/routers"
)

func main() {
	// 初始化
	initialize()
}

func initialize() {
	// 初始化goth的providers
	providers.Init()
	// 初始化gin的routers
	routers.Init()
}
