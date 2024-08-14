package account

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"user/api"
	publicProto "user/api/qvbilam/public/v1"
	proto "user/api/qvbilam/user/v1"
	"user/enum"
	"user/global"
	"user/resource"
	"user/utils"
	"user/validate"
)

func Register(ctx *gin.Context) {
	// 获取全局span
	globalSpan, _ := ctx.Get("span")
	parentSpan := globalSpan.(opentracing.Span)
	opentracing.ContextWithSpan(context.Background(), parentSpan.(opentracing.Span))
	// 将span 注入到 gin.Context 中
	context.WithValue(context.Background(), "ginContext", ctx)

	validateParamsSpan := opentracing.GlobalTracer().StartSpan("startValidateParams", opentracing.ChildOf(parentSpan.Context()))
	request := validate.CreateValidate{}
	if err := ctx.Bind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		validateParamsSpan.Finish()
		return
	}
	validateParamsSpan.Finish()

	deviceName, _ := ctx.Get("deviceName")
	deviceVersion, _ := ctx.Get("deviceVersion")
	deviceOS, _ := ctx.Get("deviceOS")
	device := proto.DeviceRequest{
		Version: deviceVersion.(string),
		Client:  deviceName.(string),
		Device:  deviceOS.(string),
	}

	sentUserServerSpan := opentracing.GlobalTracer().StartSpan("sentUserServer", opentracing.ChildOf(parentSpan.Context()))
	_, err := global.AccountServerClient.Create(ctx, &proto.UpdateAccountRequest{
		Mobile:   request.Mobile,
		Email:    request.Email,
		Password: request.Password,
		Ip:       api.GetClientIP(ctx),
		Device:   &device,
	})
	sentUserServerSpan.Finish()
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}

func Login(ctx *gin.Context) {
	request := validate.LoginValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}
	passwordMethod := map[string]string{
		enum.LoginMethodMobile:   enum.LoginMethodMobile,
		enum.LoginMethodUserName: enum.LoginMethodUserName,
		enum.LoginMethodEmail:    enum.LoginMethodEmail,
	}

	var entity *proto.AccountResponse
	var err error

	deviceName, _ := ctx.Get("deviceName")
	deviceVersion, _ := ctx.Get("deviceVersion")
	deviceOS, _ := ctx.Get("deviceOS")
	device := proto.DeviceRequest{
		Version: deviceVersion.(string),
		Client:  deviceName.(string),
		Device:  deviceOS.(string),
	}

	if _, ok := passwordMethod[request.Method]; ok {
		entity, err = global.AccountServerClient.LoginPassword(context.Background(), &proto.LoginPasswordRequest{
			Method:   request.Method,
			Username: request.Username,
			Mobile:   request.Mobile,
			Email:    request.Email,
			Password: request.Password,
			Ip:       api.GetClientIP(ctx),
			Device:   &device,
		})
	} else if request.Method == enum.LoginMethodSms {
		// 验证验证码
		if _, err := global.PublicSmsServerClient.CheckLogin(context.Background(), &publicProto.CheckSmsRequest{
			Mobile: request.Mobile,
			Code:   request.Code,
		}); err != nil {
			api.HandleGrpcErrorToHttp(ctx, err)
			return
		}

		// 验证码登陆
		entity, err = global.AccountServerClient.LoginMobile(context.Background(), &proto.LoginMobileRequest{
			Mobile: request.Mobile,
			Ip:     api.GetClientIP(ctx),
			Device: &device,
		})
	} else {
		api.ErrorUnprocessableEntity(ctx, gin.H{"method": "Unprocessable method"})
		return
	}

	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	token := entity.Token

	r := resource.LoginResource{}
	api.SuccessNotMessage(ctx, r.Resource(entity, token))
}

func LoginPlatform(ctx *gin.Context) {
	request := validate.LoginPlatformValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	allowType := []interface{}{enum.LoginMethodPlatformQQ, enum.LoginMethodPlatformWechat, enum.LoginMethodPlatformWGitHub}
	if utils.InArray(request.Type, allowType) == false {
		api.Error(ctx, "参数错误")
		return
	}

	entity, err := global.AccountServerClient.LoginPlatform(context.Background(), &proto.LoginPlatformRequest{
		Type: request.Type,
		Code: request.Code,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	token := entity.Token
	r := resource.LoginResource{}
	api.SuccessNotMessage(ctx, r.Resource(entity, token))
}

func Update(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	request := validate.UpdateAccountValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	if _, err := global.AccountServerClient.Update(context.Background(), &proto.UpdateAccountRequest{
		UserId:   userID,
		Username: request.Username,
		Mobile:   request.Mobile,
		Email:    request.Email,
		Password: request.Password,
		Ip:       api.GetClientIP(ctx),
	}); err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}
