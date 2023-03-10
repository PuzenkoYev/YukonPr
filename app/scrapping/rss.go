package scrapping

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Channel struct {
	Title string      `xml:"title"`
	Link  string      `xml:"link"`
	Desc  string      `xml:"description"`
	Items []NewsModel `xml:"item"`
}
type Enclosure struct {
	Url    string `xml:"url,attr"`
	Length int64  `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type Rss struct {
	Channel         Channel `xml:"channel"`
	ObservingStatus chan bool
}

func (n NewsModel) Stringer() {
	fmt.Println("Title:", n.Title)
	fmt.Println("Description:", n.Description)
	fmt.Println("ImageLink:", n.Image.Url)
	fmt.Println("PubDate:", n.PubDate)
	fmt.Println("FullNewsLink:", n.FullNewsLink)
}

func ParseRss(url string) *Rss {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error GET: %v\n", err)
		return &Rss{}
	}
	defer resp.Body.Close()

	rss := Rss{}
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		log.Printf("\nError Decode: %v", err)
		return &Rss{}
	} else {
		log.Printf("\nParsed succesful from :%v", url)
	}

	return &rss
}
