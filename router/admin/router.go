package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
	admin "tower/admin/controller"
	"tower/library/easycasbin"
	"tower/library/session"
	"tower/middle"
)

/**
Admin 路由
*/
func InitRouter(app *iris.Application) {

	app.Use(session.NewSessionStore())

	app.PartyFunc("/admin", func(p router.Party) {
		p.Get("/", admin.Main)
	})
	//
	// Casbin
	Egor, err := easycasbin.InitAdapter()
	if err != nil {
		panic(err)
	}
	// 使用中间件认证
	ntc := app.Party("/admin")
	{
		ntc.Use(middle.AuthAdmin(Egor, easycasbin.NotCheck("/admin/login", "/admin/logout")))
		mvc.New(ntc.Party("/user")).Handle(admin.NewManager())
	}

}
