package skills

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Skills []Skill

func unmarshalSkills(data []byte) (Skills, error) {
	var r Skills
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Skills) marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Skill struct {
	Date        string `json:"date"`
	Company     string `json:"company"`
	Link        string `json:"link"`
	LinkDisplay string `json:"linkDisplay"`
	JobTitle    string `json:"jobTitle"`
	Descr       string `json:"descr"`
}

func GetSkills(c *gin.Context) Skills {
	resp, err := http.Get("https://my-json-server.typicode.com/dbbd59/json_server/jobs")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	s, _ := unmarshalSkills(body)
	return s
}
