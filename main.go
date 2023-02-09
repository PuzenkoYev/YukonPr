package main

import (
	"yukonpr/app"
	"yukonpr/configs"
)

func main() {
	config := configs.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
