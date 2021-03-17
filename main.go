package main

import (
	"tower/boot"
	"tower/boot/router"
	"fmt"
)

func main() {
	app := router.InitRouter()
	boot.Run(app)
}
