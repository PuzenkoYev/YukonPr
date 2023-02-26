package app

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"time"
	"yukonpr/app/dbs"
	"yukonpr/app/handler"
	"yukonpr/app/models"
	"yukonpr/app/scrapping"
	"yukonpr/configs"
)

type App struct {
	Router    *mux.Router
	DB        *gorm.DB
	Scrapping *scrapping.Rss
}

func (a *App) Initialize(config *configs.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = models.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) StartObserving(url string) {
	a.Scrapping = &scrapping.Rss{}
	a.Scrapping.ObservingStatus = make(chan bool, 1)
	//defer close(a.Scrapping.ObservingStatus)
	go a.observe(url, a.Scrapping.ObservingStatus)

}

func (a *App) observe(url string, c chan bool) {
	duration := time.Minute
	defer close(c)
	for len(c) < 1 {
		a.Scrapping = scrapping.ParseRss(url)
		for _, item := range a.Scrapping.Channel.Items {
			dbs.AddNews(*a.DB, scrapping.ToNewsModel(item))
		}
		fmt.Printf("\nIt's time to sleep. Wake up at %v", time.Now().Add(duration))
		time.Sleep(duration)
	}
}

func (a *App) Run(host string) {
	fmt.Println("Running server...")
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) setRouters() {
	a.Get("/news", a.handleRequest(handler.GetFullNews))
	a.Get("/news/search", a.handleRequest(handler.GetListOfShortNews))
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}
