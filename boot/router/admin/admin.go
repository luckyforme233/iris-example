package admin

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
	admin "tower/app/handlers/admin/controller"
	"tower/app/middle"
	"tower/library/easycasbin"
	"tower/library/session"
)

/**
Admin 路由
*/
func InitRouter(app *iris.Application) {
	// 使用SESSION
	app.Use(session.NewSessionStore())
	// 使用VIEW模板
	app.RegisterView(iris.HTML("app/views", ".html"))
	fmt.Println("初始化路由")
	// 免登陆的路由
	app.PartyFunc("/admin", func(p router.Party) {
		p.Get("/", admin.Main)
		p.Get("/login", admin.Login)
		p.Post("/login", admin.Login)
	})

	// 使用中间件认证
	ntc := app.Party("/admin")
	{
		ntc.Use(middle.AuthAdmin(easycasbin.NotCheck("/admin/login", "/admin/logout")))
		mvc.New(ntc.Party("/user")).Handle(admin.NewManager())
	}

}
