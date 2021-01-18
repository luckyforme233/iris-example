package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	recover2 "github.com/kataras/iris/v12/middleware/recover"
	"github.com/spf13/viper"
	"time"
	"tower/library/config"
	"tower/library/databases"
	"tower/router/admin"
	"tower/router/api"
)

func main() {

	erre := config.Init("./conf/app.yaml")
	if erre != nil {
		panic(erre)
	}
	app := iris.New()
	// 设置日志级别
	app.Logger().SetLevel(viper.GetString("runmode"))
	// 初始化DB
	databases.InitDB()
	// 重启
	app.Use(recover2.New())
	// API 路由
	api.InitRouter(app)
	// admin 路由
	admin.InitRouter(app)

	// 是否开启TLS
	tls := viper.GetBool("tls")
	var err error
	if tls == false {
		err = app.Run(
			iris.Addr(fmt.Sprintf("%s", viper.GetString("addr"))),
			iris.WithoutServerError(iris.ErrServerClosed),
			iris.WithOptimizations,
			iris.WithTimeFormat(time.RFC3339),
		)
	} else {
		host := fmt.Sprintf("%s:%d", viper.GetString("host"), 443)
		err = app.Run(iris.TLS(host, viper.GetString("cert"), viper.GetString("key")))
	}
	if err != nil {
		fmt.Println(err)
	}
}
