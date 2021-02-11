package jsonserver_service

import (
	jsonServer "goRestServices/pkg/json-server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetJobs(c *gin.Context) {
	j := jsonServer.GetJobs(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": j})
}

func GetSkills(c *gin.Context) {
	s := jsonServer.GetSkills(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": s})
}
