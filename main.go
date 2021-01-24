package main

import (
	"tower/boot"
	"tower/boot/router"
)

func main() {
	app := router.InitRouter()
	boot.Run(app)
}
