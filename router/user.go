package router

import (
	"github.com/gin-gonic/gin"
	"user/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	VideoRouter := Router.Group("user")
	{
		VideoRouter.GET("", api.List)
		VideoRouter.GET("/:id", api.Detail)
		VideoRouter.POST("", api.Create)
		VideoRouter.PUT("/:id", api.Update)
	}
}
