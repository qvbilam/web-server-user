package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user/api/account"
	"user/api/user"
	"user/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	ServerRouter := Router.Group("user")
	{
		// 直接响应
		baseRouter := ServerRouter.Group("").Use(middleware.Cors())
		{
			baseRouter.GET("ping", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"msg": "pong",
				})
			})
			baseRouter.GET("version", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"msg": "v1.1.0",
				})
			})
		}

		// 需要登陆
		userAuthRouter := ServerRouter.Group("").Use(middleware.Cors()).Use(middleware.Auth()).Use(middleware.Trace())
		{
			userAuthRouter.GET("/search", user.Search)
			userAuthRouter.GET("/homepage/:id", user.Detail)
			userAuthRouter.PUT("/homepage", user.Update)
			userAuthRouter.POST("/account/update", account.Update)
		}

		// 不需要登陆
		accountRouter := ServerRouter.Group("").Use(middleware.Cors()).Use(middleware.Trace())
		{
			accountRouter.POST("/account/register", account.Register)
			accountRouter.POST("/account/login", account.Login)
			accountRouter.POST("/account/login/platform", account.LoginPlatform)
			accountRouter.POST("/account/logout", account.Login).Use(middleware.Auth())
		}
	}
}
