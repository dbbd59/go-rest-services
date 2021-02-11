package gh_trends_service

import (
	ghTrends "goRestServices/pkg/gh-trends"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GhTrends(c *gin.Context) {
	ght := ghTrends.GetGhTrends(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": ght})
}
