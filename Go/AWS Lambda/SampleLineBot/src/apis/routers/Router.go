package routers

import (
	"apis/controllers"
	"core/bot_handlers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) *gin.Engine {

	r.POST("/callback", bot_handlers.SampleBotHandler)
	r.GET("/callback", bot_handlers.SampleBotHandler)

	r.GET("/sample/getSomething", controllers.GetSomething)
	r.GET("/sample/doSomething", controllers.DoSomething)

	return r
}