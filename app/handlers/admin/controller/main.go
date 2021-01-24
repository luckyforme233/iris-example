package admin

import (
	"github.com/kataras/iris/v12"
	"tower/app/services"
)

func Main(ctx iris.Context) {

	ctx.View("login.html")
}

func Login(ctx iris.Context) {
	username := ctx.PostValueDefault("username", "123")
	password := ctx.PostValueDefault("password", "123")
	addr := ctx.RemoteAddr()
	auth := services.AdminAuth{}
	login := auth.Login(username, password, addr)
	ctx.JSON(login)
}
