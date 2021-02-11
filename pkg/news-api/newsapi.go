package newsapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func unmarshalNews(data []byte) (News, error) {
	var r News
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *News) marshal() ([]byte, error) {
	return json.Marshal(r)
}

type News struct {
	Status       string    `json:"status"`
	TotalResults int64     `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Article struct {
	Source      Source  `json:"source"`
	Author      *string `json:"author"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	URL         string  `json:"url"`
	URLToImage  *string `json:"urlToImage"`
	PublishedAt string  `json:"publishedAt"`
	Content     *string `json:"content"`
}

type Source struct {
	ID   *string `json:"id"`
	Name string  `json:"name"`
}

func GetNews(c *gin.Context) News {
	resp, err := http.Get("https://newsapi.org/v2/everything?q=programming&apiKey=c66c291990c048ff842cacb5aba3f53f")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	n, _ := unmarshalNews(body)
	return n
}
