package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user/middleware"
	userRouter "user/router"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors()).Use(middleware.UserAgent())
	router.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	router.GET("version", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "v1.0.0",
		})
	})
	apiRouter := router.Group("")

	// 初始化基础组建路由
	userRouter.InitUserRouter(apiRouter)
	userRouter.InitAccountRouter(apiRouter)

	return router
}
