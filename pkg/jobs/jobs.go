package jobs

import (
	"goRestServices/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindJobs(c *gin.Context) []models.Job {
	var Jobs []models.Job
	models.DB.Find(&Jobs)

	return Jobs
}

func CreateJob(c *gin.Context) models.Job {
	var input models.CreateJobInput
	if err := c.ShouldBindJSON(&input); err != nil {
		return models.Job{}
	}

	Job := models.Job{Date: input.Date, Company: input.Company, Link: input.Link, LinkDisplay: input.LinkDisplay, JobTitle: input.JobTitle, Descr: input.Descr}
	models.DB.Create(&Job)

	return Job
}

func FindJob(c *gin.Context) models.Job {
	var Job models.Job

	if err := models.DB.Where("id = ?", c.Param("id")).First(&Job).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return models.Job{}
	}

	return Job
}

func UpdateJob(c *gin.Context) models.Job {
	var Job models.Job
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Job).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return models.Job{}
	}

	var input models.UpdateJobInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return models.Job{}
	}

	models.DB.Model(&Job).Updates(input)

	return Job
}

func DeleteJob(c *gin.Context) bool {
	var Job models.Job
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Job).Error; err != nil {
		return false
	}

	models.DB.Delete(&Job)

	return true
}
