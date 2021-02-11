package news_api_service

import (
	newsapi "goRestServices/pkg/news-api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNews(c *gin.Context) {
	n := newsapi.GetNews(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": n})
}
