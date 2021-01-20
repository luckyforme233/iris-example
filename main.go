package main

import (
	"tower/boot"
	"tower/boot/router"
)

func main() {
	// 初始化HTTP 服务
	app := boot.Run()
	// 初始化路由
	router.InitRouter(app)
}
