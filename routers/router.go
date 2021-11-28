package routers

import (
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("ping", func(context *gin.Context) {
			context.JSON(200,gin.H{
				"message": "pong",
			})
		})
	}
	r.Run(utils.HttpPort)
}
