package admin

import (
	"github.com/kataras/iris/v12"
)

func Main(ctx iris.Context) {

	ctx.View("login.html")
}

func Login(ctx iris.Context) {

}
