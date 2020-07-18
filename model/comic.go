package model

import (
	"encoding/json"
	"fmt"
)

type ComicResponse struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

// FormattedDate formats individual date elements into a single string
func (cr ComicResponse) FormattedDate() string {
	return fmt.Sprintf("%s-%s-%s", cr.Day, cr.Month, cr.Year)
}

// Comic converts ComicResponse that we receive from the API to our application's output format, Comic
func (cr ComicResponse) Comic() Comic {
	return Comic{
		Title:       cr.Title,
		Number:      cr.Num,
		Date:        cr.FormattedDate(),
		Description: cr.Alt,
		Image:       cr.Img,
	}
}

type Comic struct {
	Title       string `json:"title"`
	Number      int    `json:"number"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

// PrettyString creates a pretty string of the Comic that we'll use as output
func (c Comic) PrettyPrint() string {
	p := fmt.Sprintf("Title: %s\nComic No: %s\nDate: %s\nDescription: %s\nImage: %s\n",
		cr.Title, cr.Number, cr.Date, cr.Description, cr.Image)
	return p
}

// JSON converts the Comic struct to JSON, we'll use the JSON string as output
func (c Commic) JSON() string {
	cJSON, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(cJSON)
}
