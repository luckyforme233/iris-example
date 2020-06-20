package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
)

/**
Admin 路由
*/
func InitRouter(app *iris.Application) {
	app.PartyFunc("/admin", func(p router.Party) {
		p.Get("/", func(ctx context.Context) {
			ctx.JSON(map[string]string{"hello": "1231231"})
		})
	})
}
