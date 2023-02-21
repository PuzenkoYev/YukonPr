package scrapping

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
	"yukonpr/app/models"
)

type Channel struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Desc  string `xml:"description"`
	Items []News `xml:"item"`
}
type News struct {
	Title        string    `xml:"title"`
	Description  string    `xml:"description"`
	Image        Enclosure `xml:"enclosure"`
	PubDate      string    `xml:"pubDate"`
	FullNewsLink string    `xml:"link"`
}
type Enclosure struct {
	Url    string `xml:"url,attr"`
	Length int64  `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type Rss struct {
	Channel Channel `xml:"channel"`
}

func (n News) Stringer() {
	fmt.Println("Title:", n.Title)
	fmt.Println("Description:", n.Description)
	fmt.Println("ImageLink:", n.Image.Url)
	fmt.Println("PubDate:", n.PubDate)
	fmt.Println("FullNewsLink:", n.FullNewsLink)
}
func ParseRss(url string) *Rss {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error GET: %v\n", err)
		return &Rss{}
	}
	defer resp.Body.Close()

	rss := Rss{}

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		fmt.Printf("\nError Decode: %v", err)
		return &Rss{}
	} else {
		fmt.Printf("\nParsed succesful from :%v", url)
	}

	return &rss
}

// ToNewsModel convert scrapping.News to models.NewsModel
func ToNewsModel(oldModel News) models.NewsModel {
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
