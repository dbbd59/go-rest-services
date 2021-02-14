package jobs_service

import (
	"goRestServices/pkg/jobs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindJobs(c *gin.Context) {
	jobs := jobs.FindJobs(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": jobs})
}

func CreateJob(c *gin.Context) {
	job := jobs.CreateJob(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": job})
}

func FindJob(c *gin.Context) {
	job := jobs.FindJob(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": job})
}

func UpdateJob(c *gin.Context) {
	job := jobs.UpdateJob(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": job})
}

func DeleteJob(c *gin.Context) {
	r := jobs.DeleteJob(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": r})
}
