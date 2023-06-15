package routes

import (
	"mygin/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		//user模块路由

		//article模块路由

		//category模块路由
	}
	r.Run(utils.HttpPort)
}
