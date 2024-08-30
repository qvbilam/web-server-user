package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"user/global"
)

func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 配置
		cfg := config.Configuration{
			Sampler: &config.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,                          // 始终做出相同决定的采样器类型。
				Param: float64(global.ServerConfig.JaegerConfig.Output), // 0 不采样, 1全部采样
			},
			Reporter: &config.ReporterConfig{
				LogSpans:           global.ServerConfig.JaegerConfig.IsLog, // 是否打印日志
				LocalAgentHostPort: fmt.Sprintf("%s:%s", global.ServerConfig.JaegerConfig.Host, global.ServerConfig.JaegerConfig.Port),
			},
			ServiceName: global.ServerConfig.JaegerConfig.Server,
		}
		// 生成链路Tracer
		tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
		if err != nil {
			panic(any(err))
		}
		// 关闭链路
		defer func(closer io.Closer) {
			err := closer.Close()
			if err != nil {
				panic(any(err))
			}
		}(closer)

		// 设置全局 tracer
		opentracing.SetGlobalTracer(tracer)
		span := tracer.StartSpan(ctx.Request.URL.Path) // 设置 span 为请求url
		defer span.Finish()                            // span 完成

		// 定义上下文，用于 child span 获取 父级
		ctx.Set("tracer", tracer)
		ctx.Set("span", span)
		ctx.Next()
	}
}
