package routers

import (
	"goRestServices/service/hn_service"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/hn/:story", hn_service.HackerNews)
	}
	return r
}
