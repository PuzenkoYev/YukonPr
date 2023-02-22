package main

import (
	"fmt"
	"time"
	"yukonpr/app"
	"yukonpr/configs"
)

func main() {
	config := configs.GetConfig()

	app := &app.App{}
	r := time.Now().Add(time.Hour)
	fmt.Println(r)
	r2 := time.Now().Add(-2 * time.Hour)
	fmt.Println(r2)
	r3 := time.Now().Add(-10 * time.Hour)
	fmt.Println(r3)

	app.Initialize(config)
	app.Run(":3000")

	//fmt.Println(db.SelectByWordInTitleOrText(*app.DB, "Украина"))
	app.StartObserving("https://rss.unian.net/site/news_rus.rss")

}
