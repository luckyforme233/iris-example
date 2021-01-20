package admin

import (
	"github.com/kataras/iris/v12"
)

func Main(ctx iris.Context) {
	//ctx.JSON(iris.Map{"test" : 1})

	ctx.View("login.html")
}
