package main

import (
	"fmt"
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

	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}

	server.ListenAndServe()

}
