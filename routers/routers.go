package routers

import (
	"goRestServices/service/books_service"
	gh_trends_service "goRestServices/service/gh_service"
	"goRestServices/service/hn_service"
	"goRestServices/service/jobs_service"
	news_api_service "goRestServices/service/news_service"
	"goRestServices/service/skills_service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(cors.Default())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/hn/:story", hn_service.HackerNews)
		apiv1.GET("/ghtrends", gh_trends_service.GhTrends)
		apiv1.GET("/news", news_api_service.GetNews)

		books := apiv1.Group("/books")
		{
			books.GET("", books_service.FindBooks)
			books.POST("", books_service.CreateBook)
			books.PATCH("/:id", books_service.UpdateBook)
			books.DELETE("/:id", books_service.DeleteBook)
		}

		jobs := apiv1.Group("/jobs")
		{
			jobs.GET("", jobs_service.FindJobs)
			jobs.POST("", jobs_service.CreateJob)
			jobs.PATCH("/:id", jobs_service.UpdateJob)
			jobs.DELETE("/:id", jobs_service.DeleteJob)
		}

		skills := apiv1.Group("/skills")
		{
			skills.GET("", skills_service.FindSkills)
			skills.POST("", skills_service.CreateSkill)
			skills.PATCH("/:id", skills_service.UpdateSkill)
			skills.DELETE("/:id", skills_service.DeleteSkill)
		}
	}
	return r
}
