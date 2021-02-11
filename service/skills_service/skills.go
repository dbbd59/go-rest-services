package skills_service

import (
	"goRestServices/pkg/skills"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GhSkills(c *gin.Context) {
	s := skills.GetSkills(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": s})
}
