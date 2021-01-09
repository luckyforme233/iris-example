package admin

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

func Main(ctx iris.Context) {
	fmt.Println(ctx.GetCurrentRoute().Path())
	ctx.JSON(map[string]string{"ddd": "222"})
}
