package handler

import (
	"time"
	"yukonpr/app/models"
)

type NewsModel struct {
	Title       string    `json:"Title"`
	PubTime     time.Time `json:"Time"`
	Description string    `json:"Description"`
	ImageUrl    string    `json:"ImageUrl"`
}

type ShortNewsModel struct {
	Title   string    `json:"Title"`
	PubTime time.Time `json:"Time"`
}

func ToHandlerNewsModel(oldModel models.NewsModel) NewsModel {
	var newModel NewsModel
	newModel.Title = oldModel.Title
	newModel.Description = oldModel.Description
	newModel.ImageUrl = oldModel.ImageUrl
	newModel.PubTime = oldModel.PubTime

	return newModel
}
func ToHandlerShortNewsModel(oldModel models.NewsModel) ShortNewsModel {
	var newModel ShortNewsModel
	newModel.Title = oldModel.Title
	newModel.PubTime = oldModel.PubTime

	return newModel
}
