package handler

import (
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
	"yukonpr/app/dbs"
)

// GetFullNews /news?id - GET
func GetFullNews(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println(err)
		respondJSON(w, http.StatusBadRequest, []NewsModel{})
		return
	}

	if dbs.ContainsById(*db, id) == false {
		log.Printf("News not found. Title {%d}", id)
		respondJSON(w, http.StatusBadRequest, NewsModel{})
		return
	} else {
		news := ToHandlerNewsModel(dbs.SelectById(*db, id))
		respondJSON(w, http.StatusOK, news)
	}
}

// GetListOfShortNews /news/search - GET
func GetListOfShortNews(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	//Is the request .../search?contains
	word := r.URL.Query().Get("contains")
	if word != "" {
		GetContainedWordShortNews(db, w, word)
		return
	} else {
		//The request is .../search?from&to
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
	return
}
