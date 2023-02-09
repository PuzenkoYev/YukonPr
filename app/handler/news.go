package handler

import (
	"encoding/json"
	"net/http"
	"time"
	"yukonpr/app/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func AddNews(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	news := models.NewsModel{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&news); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&news).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, news)
}
func GetAllNews(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	news := []models.NewsModel{}
	db.Find(&news)
	respondJSON(w, http.StatusOK, news)
}

func GetNews(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["Title"]
	project := getNewsOr404(db, title, w, r)
	if project == nil {
		return
	}
	respondJSON(w, http.StatusOK, project)
}

func getNewsOr404(db *gorm.DB, title string, w http.ResponseWriter, r *http.Request) *models.NewsModel {
	news := models.NewsModel{}
	if err := db.First(&news, models.NewsModel{Title: title}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &news
}

// TODO: Check work it or not
func GetNewsByTags(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var news []models.NewsModel
	var tags []models.Tags

	tagNames := r.URL.Query()["tag"]

	// Query the database for all news with the given tags
	db.Where("name IN (?)", tagNames).Find(&tags)

	// Get all news with the given tags
	for _, tag := range tags {
		db.Where("id = ?", tag.NewsId).Find(&news)
	}

	respondJSON(w, http.StatusOK, news)
}
func GetNewsBetweenTimes(from time.Time, to time.Time) ([]models.NewsModel, error)
func GetNewsContainsWord(word string) ([]models.NewsModel, error)
