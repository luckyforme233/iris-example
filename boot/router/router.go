package router

import (
	"github.com/kataras/iris/v12"
	"tower/boot/router/admin"
	"tower/boot/router/api"
)

func InitRouter(app *iris.Application) {
	// API 路由
	api.InitRouter(app)
	// admin 路由
	admin.InitRouter(app)
}
