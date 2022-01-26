package main

import (
	"lib_goth_gin/providers"
	"lib_goth_gin/routers"
)

func init() {
	// 初始化goth的providers
	providers.Init()
	// 初始化gin的routers
	routers.Init()
}

func main() {

}
