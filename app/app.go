package app

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"time"
	"yukonpr/app/db"
	"yukonpr/app/scrapping"
)

type App struct {
	Router    *mux.Router
	DB        *gorm.DB
	Scrapping *scrapping.Rss
}

//func (a *App) Initialize(config *configs.Config) {
//	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
//		config.DB.Username,
//		config.DB.Password,
//		config.DB.Host,
//		config.DB.Port,
//		config.DB.Name,
//		config.DB.Charset)
//
//	db, err := gorm.Open(config.DB.Dialect, dbURI)
//	if err != nil {
//		log.Fatal("Could not connect database")
//	}
//
//	a.DB = models.DBMigrate(db)
//	a.Router = mux.NewRouter()
//	a.setRouters()
//}

func (a *App) StartObserving(url string) {
	doObserve := make(chan bool, 1)
	defer close(doObserve)
	go a.observe(url, doObserve)

	for true {
		var str string
		fmt.Scanln(&str)
		if str == "stop" {
			doObserve <- false
			break
		}
	}
}

func (a *App) observe(url string, c chan bool) {
	duration := 5 * time.Minute
	for len(c) < 1 {
		a.Scrapping = scrapping.ParseRss(url)
		for _, item := range a.Scrapping.Channel.Items {
			db.AddNews(*a.DB, scrapping.ToNewsModel(item))
		}
		fmt.Printf("\nIt's time to sleep. Wake up at %v", time.Now().Add(duration))
		time.Sleep(duration)
	}
}
