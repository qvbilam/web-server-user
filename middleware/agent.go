package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mssola/useragent"
)

// UserAgent 解析user-agent
func UserAgent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ua := useragent.New(ctx.GetHeader("user-agent"))
		name, version := ua.Browser()
		ctx.Set("deviceName", name)
		ctx.Set("deviceVersion", version)
		ctx.Set("deviceOS", ua.OS())

		// 继续执行
		ctx.Next()
	}
}
