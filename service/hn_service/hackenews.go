package hn_service

import (
	hn "goRestServices/pkg/hacker-news"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HackerNews(c *gin.Context) {
	story := c.Param("story")

	hnr := hn.GetHackerNews(c, story)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": hnr})
}
