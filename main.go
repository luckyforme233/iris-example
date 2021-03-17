package main

import (
	"tower/boot"
	"tower/boot/router"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	app := router.InitRouter()
	boot.Run(app)
}
