package initialize

import (
	"github.com/gin-gonic/gin"
	"user/middleware"
	userRouter "user/router"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors()).Use(middleware.UserAgent())
	apiRouter := router.Group("")

	// 初始化基础组建路由
	userRouter.InitUserRouter(apiRouter)
	return router
}
