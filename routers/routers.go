package routers

import (
	gh_trends_service "goRestServices/service/gh_service"
	"goRestServices/service/hn_service"
	"goRestServices/service/jsonserver_service"
	news_api_service "goRestServices/service/news_service"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/hn/:story", hn_service.HackerNews)
		apiv1.GET("/ghtrends", gh_trends_service.GhTrends)
		apiv1.GET("/jobs", jsonserver_service.GetJobs)
		apiv1.GET("/skills", jsonserver_service.GetSkills)
		apiv1.GET("/news", news_api_service.GetNews)
	}
	return r
}
