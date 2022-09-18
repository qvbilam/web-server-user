package router

import (
	"github.com/gin-gonic/gin"
	"user/api"
	"user/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	ServerRouter := Router.Group("user")
	{
		userAuthRouter := ServerRouter.Group("").Use(middleware.Cors()).Use(middleware.Auth())
		{
			userAuthRouter.GET("", api.Search) // todo
			userAuthRouter.GET("/search", api.Search)
			userAuthRouter.GET("/:id", api.Detail)
			userAuthRouter.PUT("/:id", api.Update)
			userAuthRouter.POST("/logout", api.Logout)
		}

		userRouter := ServerRouter.Group("").Use(middleware.Cors())
		{
			userRouter.POST("/register", api.Register)
			userRouter.POST("/login", api.Login)
		}
	}
}
