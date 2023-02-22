package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
	"yukonpr/app/dbs"
)

// GetFullNews news/{title} - GET
func GetFullNews(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	if title == "" {
		println("Title is empty")
		respondJSON(w, http.StatusBadRequest, NewsModel{})
		return
	}
	if dbs.ContainsByTitle(*db, title) == false {
		fmt.Printf("Title was not found. Title {%v}", title)
		respondJSON(w, http.StatusBadRequest, NewsModel{})
		return
	} else {
		news := ToHandlerNewsModel(dbs.SelectByTitle(*db, title))
		respondJSON(w, http.StatusOK, news)
	}
}

// GetListShortNewsBetweenTimes news/{{from},{to}} - GET
func GetListShortNewsBetweenTimes(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	layout := "2006-01-02 15:04:05"

	temp := vars["from"]
	from, err := time.Parse(layout, temp)
	if err != nil {
		println(err)
		respondJSON(w, http.StatusBadRequest, []ShortNewsModel{})
		return
	}
	to, err := time.Parse(layout, vars["to"])
	if err != nil {
		println(err)
		respondJSON(w, http.StatusBadRequest, []ShortNewsModel{})
		return
	}

	news := []ShortNewsModel{}
	selectedItems := dbs.SelectInTimeRange(*db, from, to)
	for _, item := range selectedItems {
		news = append(news, ToHandlerShortNewsModel(item))
	}

	respondJSON(w, http.StatusOK, news)
}

//
//import (
//	"encoding/json"
//	"log"
//	"net/http"
//	"yukonpr/app/models"
//
//	"github.com/gorilla/mux"
//	"github.com/jinzhu/gorm"
//)
//
//func CreateNews(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
//	defer r.Body.Close()
//	news := models.NewsModel{}
//
//	decoder := json.NewDecoder(r.Body)
//	if err := decoder.Decode(&news); err != nil {
//		respondError(w, http.StatusBadRequest, err.Error())
//		return
//	}
//
//	if err := db.Save(&news).Error; err != nil {
//		respondError(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//	log.Print("Created", news)
//
//	respondJSON(w, http.StatusCreated, news)
//}
//func GetAllNews(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
//	news := []models.NewsModel{}
//	db.Find(&news)
//
//	respondJSON(w, http.StatusOK, news)
//}
//
//func GetNews(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	title := vars["Title"]
//	project := getNewsOr404(db, title, w, r)
//	if project == nil {
//		return
//	}
//
//	respondJSON(w, http.StatusOK, project)
//}
//
////TODO: Check work it or not
//func GetNewsByTags(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
//	var news []models.NewsModel
//	var tags []models.Tags
//
//	tagNames := r.URL.Query()["tag"]
//	db.Where("Name IN (?)", tagNames).Find(&tags)
//	for _, tag := range tags {
//		db.Where("ID = ?", tag.NewsId).Find(&news)
//	}
//
//	respondJSON(w, http.StatusOK, news)
//}
//
//// TODO: Check work it or not
//func GetNewsBetweenTimes(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
//	var news []models.NewsModel
//
//	startTime := r.URL.Query().Get("startTime")
//	endTime := r.URL.Query().Get("endTime")
//	db.Where("Time BETWEEN ? AND ?", startTime, endTime).Find(&news)
//
//	respondJSON(w, http.StatusOK, news)
//}
//
//// TODO: Check work it or not
//func GetNewsContainsWord(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
//	word := r.URL.Query().Get("word")
//	var news []models.NewsModel
//	db.Where("Title LIKE ? OR Description LIKE ?", word, word).Find(&news)
//	json.NewEncoder(w).Encode(news)
//}
//
//func getNewsOr404(db *gorm.DB, title string, w http.ResponseWriter, r *http.Request) *models.NewsModel {
//	news := models.NewsModel{}
//
//	if err := db.First(&news, models.NewsModel{Title: title}).Error; err != nil {
//		respondError(w, http.StatusNotFound, err.Error())
//		return nil
//	}
//
//	return &news
//}
