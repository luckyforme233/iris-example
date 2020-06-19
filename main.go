package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
	"tower/library/config"
	"tower/router/admin"
	"tower/router/api"
)

func main() {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	path = path[:index]
	erre := config.Init("./conf/app.yaml")
	if erre != nil {
		panic(erre)
	}
	fmt.Println(viper.GetString("host"))
	app := iris.New()
	// 设置日志级别
	app.Logger().SetLevel(viper.GetString("debug"))

	module := viper.GetStringMap("module")
	fmt.Println(module)

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
