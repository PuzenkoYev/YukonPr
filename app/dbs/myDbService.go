package dbs

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
	"yukonpr/app/models"
)

func AddNews(db gorm.DB, model models.NewsModel) bool {
	if ContainsByTitle(db, model.Title) == true {
		//fmt.Printf("\nAlready exist: Title %v",model.Title)
		return false
	}
	if err := db.Save(&model).Error; err != nil {
		log.Printf("\n%v", err.Error())

		return false
	}
	log.Printf("\nAdd new News: {%v}  {%v}", model.PubTime, model.Title)
	return true
}

func ContainsById(db gorm.DB, id int) bool {
	var item models.NewsModel
	if err := db.Where("id = ?", id).Find(&item).Error; err != nil {
		log.Printf("\n%v", err.Error())
		return false
	}
	return true
}

func ContainsByTitle(db gorm.DB, title string) bool {
	var item models.NewsModel
	if err := db.Where("title = ?", title).Find(&item).Error; err != nil {
		log.Printf("\n%v", err.Error())
		return false
	}
	return true
}

func SelectById(db gorm.DB, id int) models.NewsModel {
	var news models.NewsModel
	if ContainsById(db, id) == true {
		db.Where("id = ? ", id).Find(&news)
		return news
	} else {
		log.Printf("\nDidn't found model with id: %d", id)
		return news
	}
}

func SelectInTimeRange(db gorm.DB, from time.Time, to time.Time) []models.NewsModel {
	var news []models.NewsModel
	db.Where("pub_time BETWEEN ? AND ?", from, to).Find(&news)
	return news
}

func SelectByWordInTitleOrText(db gorm.DB, word string) []models.NewsModel {
	var news []models.NewsModel
	db.Where("title LIKE ? OR description LIKE ?", "%"+word+"%", "%"+word+"%").Find(&news)
	return news
}
