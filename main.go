package main

import (
	"yukonpr/app"
	"yukonpr/configs"
)

func main() {
	config := configs.GetConfig()

	app := &app.App{}

	//db ini
	app.Initialize(config)

	//start parsing
	app.StartObserving("https://rss.unian.net/site/news_rus.rss")

	//start api server
	app.Run(":3000")
}
