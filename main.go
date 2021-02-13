package main

import (
	"fmt"
	"goRestServices/models"
	"goRestServices/routers"
	"goRestServices/setting"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	models.ConnectDatabase()

	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}

	server.ListenAndServe()

}
