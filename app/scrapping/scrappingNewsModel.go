package scrapping

import (
	"log"
	"time"
	"yukonpr/app/models"
)

type NewsModel struct {
	Title        string    `xml:"title"`
	Description  string    `xml:"description"`
	Image        Enclosure `xml:"enclosure"`
	PubDate      string    `xml:"pubDate"`
	FullNewsLink string    `xml:"link"`
}

// ToNewsModel convert scrapping.NewsModel to models.NewsModel
func ToNewsModel(oldModel NewsModel) models.NewsModel {
	var newModel models.NewsModel
	layout := "Mon, 02 Jan 2006 15:04:05 -0700"

	newModel.Title = oldModel.Title
	newModel.ImageUrl = oldModel.Image.Url
	newModel.Description = oldModel.Description

	t, err := time.Parse(layout, oldModel.PubDate)

	if err != nil {
		log.Println(err)
	}
	newModel.PubTime = t
	return newModel
}
