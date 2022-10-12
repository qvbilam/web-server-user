package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"user/global"
	"user/initialize"
	"user/rule"
)

func main() {
	// 初始化日志
	initialize.InitLogger()
	// 初始化配置
	initialize.InitConfig()
	// 初始化路由
	Router := initialize.InitRouters()
	// 初始化表单验证
	if err := initialize.InitValidateTran("zh"); err != nil {
		zap.S().Panic("翻译器初始化失败: ", err.Error())
	}
	// 注册自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", rule.ValidateMobile) // 注册手机号验证(自定义名, 规则函数)
		// 自定义错误
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法手机号码", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	// 初始化grpc客户端
	initialize.InitServer()

	Name := global.ServerConfig.Name
	Host := "0.0.0.0"
	Port := 9701

	// 启动服务
	go func() {
		zap.S().Infof("%s start listen: %s:%d", Name, Host, Port)

		if err := Router.Run(fmt.Sprintf(":%d", Port)); err != nil {
			zap.S().Panic("%s 服务启动失败: %s", Name, err.Error())
		}
	}()

	// 监听结束
	// 接受终止信号(优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// 服务注销
	zap.S().Info("服务注销成功")
}
