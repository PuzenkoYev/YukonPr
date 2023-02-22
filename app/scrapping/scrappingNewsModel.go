package scrapping

import (
	"fmt"
	"time"
	"yukonpr/app/models"
)

type ScrappingNewsModel struct {
	Title        string    `xml:"title"`
	Description  string    `xml:"description"`
	Image        Enclosure `xml:"enclosure"`
	PubDate      string    `xml:"pubDate"`
	FullNewsLink string    `xml:"link"`
}

// ToNewsModel convert scrapping.ScrappingNewsModel to models.NewsModel
func ToNewsModel(oldModel ScrappingNewsModel) models.NewsModel {
	var newModel models.NewsModel
	layout := "Mon, 02 Jan 2006 15:04:05 -0700"

	newModel.Title = oldModel.Title
	newModel.ImageUrl = oldModel.Image.Url
	newModel.Description = oldModel.Description

	t, err := time.Parse(layout, oldModel.PubDate)

	if err != nil {
		fmt.Println(err)
	}
	newModel.PubTime = t
	return newModel
}
