package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type NewsModel struct {
	gorm.Model
	Title string    `josn:"Title"`
	Time  time.Time `json:"Time"`
	Text  string    `json:"Text"`
	Image string    `json:"Image"`
	Tags  []Tags    `gorm:"ForeignKey:NewsId" json:"NewsId"`
}
type Tags struct {
	gorm.Model
	NewsId int    `json:"NewsId"`
	Name   string `json:"Name"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&NewsModel{}, &Tags{})
	db.Model(&Tags{}).AddForeignKey("NewsId", "NewsModel(id)", "CASCADE", "CASCADE")
	return db
}
