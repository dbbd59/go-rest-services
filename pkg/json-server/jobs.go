package jsonServer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Jobs []Job

func unmarshalJobs(data []byte) (Jobs, error) {
	var r Jobs
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Jobs) marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Job struct {
	Date        string `json:"date"`
	Company     string `json:"company"`
	Link        string `json:"link"`
	LinkDisplay string `json:"linkDisplay"`
	JobTitle    string `json:"jobTitle"`
	Descr       string `json:"descr"`
}

func GetJobs(c *gin.Context) Jobs {
	resp, err := http.Get("https://my-json-server.typicode.com/dbbd59/json_server/jobs")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	j, _ := unmarshalJobs(body)
	return j
}
