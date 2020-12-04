package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	admin "tower/admin/controller"
)

/**
Admin 路由
*/
func InitRouter(app *iris.Application) {
	//app.PartyFunc("/admin", func(p router.Party) {
	//	p.Get("/", admin.Main)
	//
	//})
	//
	mvc.New(app.Party("/admin/user")).Handle(admin.NewManager())

}
