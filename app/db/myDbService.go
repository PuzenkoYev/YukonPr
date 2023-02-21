package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
	"yukonpr/app/models"
)

func AddNews(db gorm.DB, model models.NewsModel) bool {
	if err := db.Save(&model).Error; err != nil {
		fmt.Printf(err.Error())
		return false
	}
	return true
}
func RemoveNews(db gorm.DB, id int) bool {
	if err := db.Where("id = ?", id).Delete(models.NewsModel{}).Error; err != nil {
		fmt.Printf(err.Error())
		return false
	}
	return true
}
func ContainsById(db gorm.DB, id int) bool {
	var item models.NewsModel
	if err := db.Where("id = ?", id).Find(&item).Error; err != nil {
		fmt.Printf(err.Error())
		return false
	}
	return true
}
func ContainsByTitle(db gorm.DB, title string) bool {
	var item models.NewsModel
	if err := db.Where("Title = ?", title).Find(&item).Error; err != nil {
		fmt.Printf(err.Error())
		return false
	}
	return true
}
func SelectAll(db gorm.DB) []models.NewsModel {
	var news []models.NewsModel
	db.Find(&news)
	return news
}
func SelectById(db gorm.DB, id int) models.NewsModel {
	var news models.NewsModel
	if ContainsById(db, id) == true {
		db.Where("id = ? ", id).Find(&news)
		return news
	} else {
		fmt.Printf("Didn't found model with id: %d", id)
		return news
	}
}
func SelectByTitle(db gorm.DB, title string) models.NewsModel {
	var news models.NewsModel
	if ContainsByTitle(db, title) == true {
		db.Where("Title = ? ", title).Find(&news)
		return news
	} else {
		fmt.Printf("Didn't found model with id: %s", title)
		return news
	}
}
func SelectInTimeRange(db gorm.DB, from time.Time, to time.Time) []models.NewsModel {
	var news []models.NewsModel
	db.Where("Time BETWEEN ? AND ?", from, to).Find(&news)
	return news
}
func SelectByWordInTitleOrText(db gorm.DB, word string) []models.NewsModel {
	var news []models.NewsModel
	db.Where("Title LIKE ? OR Text LIKE ?", word, word).Find(&news)
	return news
}
