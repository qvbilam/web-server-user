package router

import (
	"github.com/gin-gonic/gin"
	"user/api/user"
	"user/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	ServerRouter := Router.Group("user")
	{
		userAuthRouter := ServerRouter.Group("").Use(middleware.Cors()).Use(middleware.Auth())
		{
			userAuthRouter.GET("", user.Search) // todo
			userAuthRouter.GET("/search", user.Search)
			userAuthRouter.GET("/:id", user.Detail)
			userAuthRouter.PUT("/:id", user.Update)

		}
	}
}
