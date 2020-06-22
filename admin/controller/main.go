package admin

import (
	"github.com/kataras/iris/v12"
)

func Main(ctx iris.Context) {
	ctx.JSON(map[string]string{"ddd": "222"})
}
