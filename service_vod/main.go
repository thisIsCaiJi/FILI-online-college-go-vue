package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_vod/config"
	"github.com/thisIsCaiJi/online_college/service_vod/handlers"
)


func main() {
	app := handlers.InitRouter()
	ServerRun(app)
}

func ServerRun(app *gin.Engine) {
	port := config.Yml.ServerConf.Port
	app.Run(fmt.Sprintf(":%d",port))
}
