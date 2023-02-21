package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type NewsModel struct {
	gorm.Model
	Title       string    `json:"Title"`
	PubTime     time.Time `json:"Time"`
	Description string    `json:"Description"`
	ImageUrl    string    `json:"ImageUrl"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&NewsModel{})
	return db
}
