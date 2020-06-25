package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_edu/controller"
)

func userRouter(group *gin.RouterGroup){
	group.POST("/login",controller.Login)
	group.GET("/info",controller.Info)
	group.POST("/logout",controller.Logout)
}