package main

import (
	"fmt"
	"time"
	"yukonpr/app"
	"yukonpr/app/db"
	"yukonpr/app/scrapping"
	"yukonpr/configs"
)

func main() {
	config := configs.GetConfig()

	app := &app.App{}

	app.Initialize(config)
	//app.Run(":3000")

	duration := 5 * time.Minute
	for {
		app.Scrapping = scrapping.ParseRss("https://rss.unian.net/site/news_rus.rss")
		for _, item := range app.Scrapping.Channel.Items {
			db.AddNews(*app.DB, scrapping.ToNewsModel(item))
		}
		fmt.Printf("\nIt's time to sleep. Wake up at %v", time.Now().Add(duration))
		time.Sleep(duration)
	}
}
