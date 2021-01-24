package admin

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

func Main(ctx iris.Context) {

	ctx.View("login.html")
}

func Login(ctx iris.Context) {
	//var req = &requests.Login{}
	//err := ctx.ReadJSON(req)
	//fmt.Println(err)
	//fmt.Println(ctx.ReadForm(req))
	param := ctx.PostValueDefault("username", "123")
	passwd := ctx.PostValueDefault("password", "123")
	fmt.Println(param, passwd)
}
