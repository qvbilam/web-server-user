package router

import (
	"github.com/gin-gonic/gin"
	"user/api/account"
	"user/middleware"
)

func InitAccountRouter(Router *gin.RouterGroup) {
	ServerRouter := Router.Group("account")
	{
		accountRouter := ServerRouter.Group("").Use(middleware.Cors())
		{
			accountRouter.POST("/register", account.Register)
			accountRouter.POST("/login", account.Login)
			accountRouter.POST("/logout", account.Login)
		}
	}
}
