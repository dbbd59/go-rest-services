package main

import (
	"goRestServices/models"
	"goRestServices/routers"
	"goRestServices/setting"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	models.SetupAndConnectDatabase()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	endPoint := setting.GetEndpoint()

	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}

	server.ListenAndServe()

}
