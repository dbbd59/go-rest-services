package skills

import (
	"goRestServices/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindSkills(c *gin.Context) []models.Skill {
	var Skills []models.Skill
	models.DB.Find(&Skills)

	return Skills
}

func CreateSkill(c *gin.Context) models.Skill {
	var input models.CreateSkillInput
	if err := c.ShouldBindJSON(&input); err != nil {
		return models.Skill{}
	}

	Skill := models.Skill{Name: input.Name, Perc: input.Perc}
	models.DB.Create(&Skill)

	return Skill
}

func FindSkill(c *gin.Context) models.Skill {
	var Skill models.Skill

	if err := models.DB.Where("id = ?", c.Param("id")).First(&Skill).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return models.Skill{}
	}

	return Skill
}

func UpdateSkill(c *gin.Context) models.Skill {
	var Skill models.Skill
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Skill).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return models.Skill{}
	}

	var input models.UpdateSkillInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return models.Skill{}
	}

	models.DB.Model(&Skill).Updates(input)

	return Skill
}

func DeleteSkill(c *gin.Context) bool {
	var Skill models.Skill
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Skill).Error; err != nil {
		return false
	}

	models.DB.Delete(&Skill)

	return true
}
