package ghTrends

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GhTrends []GhTrend

type GhTrend struct {
	Author             string    `json:"author"`
	Name               string    `json:"name"`
	Avatar             string    `json:"avatar"`
	Description        *string   `json:"description"`
	URL                string    `json:"url"`
	Language           *string   `json:"language"`
	LanguageColor      *string   `json:"languageColor"`
	Stars              int64     `json:"stars"`
	Forks              int64     `json:"forks"`
	CurrentPeriodStars int64     `json:"currentPeriodStars"`
	BuiltBy            []BuiltBy `json:"builtBy"`
}

type BuiltBy struct {
	Username string `json:"username"`
	Href     string `json:"href"`
	Avatar   string `json:"avatar"`
}

func unmarshalGhTrends(data []byte) (GhTrends, error) {
	var r GhTrends
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GhTrends) marshal() ([]byte, error) {
	return json.Marshal(r)
}

func GetGhTrends(c *gin.Context) GhTrends {
	resp, err := http.Get("https://gtrend.yapie.me/repositories")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	ght, _ := unmarshalGhTrends(body)
	return ght
}
