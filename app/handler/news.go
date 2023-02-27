package handler

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
	"yukonpr/app/dbs"
)

// GetFullNews /news - GET
func GetFullNews(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	if search != "" {
		GetContainedWordShortNews(db, w, search)
		return
	}
}

// GetFullNewsById /news/{id} - GET
func GetFullNewsById(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)

	id, err := strconv.Atoi(request["id"])
	if err != nil {
		log.Println(err)
		return
	}
	news := ToHandlerNewsModel(dbs.SelectById(*db, id))
	respondJSON(w, http.StatusOK, news)
}

// GetListOfShortNews /news/search - GET
func GetListOfShortNews(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	layout := "2006-01-02 15:04:05"

	from, err := time.Parse(layout, r.URL.Query().Get("from"))
	if err != nil {
		log.Println(err)
		respondJSON(w, http.StatusBadRequest, []ShortNewsModel{})
		return
	}
	to, err := time.Parse(layout, r.URL.Query().Get("to"))
	if err != nil {
		log.Println(err)
		respondJSON(w, http.StatusBadRequest, []ShortNewsModel{})
		return
	}
	GetShortNewsBetweenTime(db, w, from, to)
}

// GetShortNewsBetweenTime /news/search?from&to - GET
func GetShortNewsBetweenTime(db *gorm.DB, w http.ResponseWriter, from time.Time, to time.Time) {
	news := []ShortNewsModel{}
	selectedItems := dbs.SelectInTimeRange(*db, from, to)
	for _, item := range selectedItems {
		news = append(news, ToHandlerShortNewsModel(item))
	}

	respondJSON(w, http.StatusOK, news)
}

// GetContainedWordShortNews /news/search?word - GET
func GetContainedWordShortNews(db *gorm.DB, w http.ResponseWriter, word string) {
	arr := dbs.SelectByWordInTitleOrText(*db, word)
	var res []ShortNewsModel
	for _, val := range arr {
		res = append(res, ToHandlerShortNewsModel(val))
	}
	respondJSON(w, http.StatusOK, res)
}
