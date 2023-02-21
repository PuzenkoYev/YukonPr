package handler

//
//import (
//	"encoding/json"
//	"net/http"
//	"strconv"
//	"yukonpr/app/models"
//
//	"github.com/gorilla/mux"
//	"github.com/jinzhu/gorm"
//)
//
//func GetAllTags(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//
//	newsTitle := vars["Title"]
//	project := getNewsOr404(db, newsTitle, w, r)
//	if project == nil {
//		return
//	}
//
//	tags := []models.Tags{}
//	if err := db.Model(&project).Related(&tags).Error; err != nil {
//		respondError(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	respondJSON(w, http.StatusOK, tags)
//}
//func CreateTags(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//
//	newsTitle := vars["title"]
//	news := getNewsOr404(db, newsTitle, w, r)
//	if news == nil {
//		return
//	}
//
//	tag := models.Tags{NewsId: int(news.ID)}
//
//	decoder := json.NewDecoder(r.Body)
//	if err := decoder.Decode(&tag); err != nil {
//		respondError(w, http.StatusBadRequest, err.Error())
//		return
//	}
//	defer r.Body.Close()
//
//	if err := db.Save(&tag).Error; err != nil {
//		respondError(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	respondJSON(w, http.StatusCreated, tag)
//}
//
//func GetTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//
//	newsTitle := vars["title"]
//	news := getNewsOr404(db, newsTitle, w, r)
//	if news == nil {
//		return
//	}
//
//	id, _ := strconv.Atoi(vars["id"])
//	tag := getTagsOr404(db, id, w, r)
//	if tag == nil {
//		return
//	}
//
//	respondJSON(w, http.StatusOK, tag)
//}
//
//func getTagsOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *models.Tags {
//	tags := models.Tags{}
//	if err := db.First(&tags, id).Error; err != nil {
//		respondError(w, http.StatusNotFound, err.Error())
//		return nil
//	}
//
//	return &tags
//}
