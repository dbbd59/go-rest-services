package skills_service

import (
	"goRestServices/pkg/skills"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindSkills(c *gin.Context) {
	skills := skills.FindSkills(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": skills})
}

func CreateSkill(c *gin.Context) {
	skill := skills.CreateSkill(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": skill})
}

func FindSkill(c *gin.Context) {
	skill := skills.FindSkill(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": skill})
}

func UpdateSkill(c *gin.Context) {
	skill := skills.UpdateSkill(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": skill})
}

func DeleteSkill(c *gin.Context) {
	r := skills.DeleteSkill(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": r})
}
