package account

import (
	"context"
	"github.com/gin-gonic/gin"
	"user/api"
	publicProto "user/api/qvbilam/public/v1"
	proto "user/api/qvbilam/user/v1"
	"user/enum"
	"user/global"
	"user/resource"
	"user/validate"
)

func Register(ctx *gin.Context) {
	request := validate.CreateValidate{}
	if err := ctx.Bind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	_, err := global.AccountServerClient.Create(context.Background(), &proto.UpdateAccountRequest{
		Mobile:   request.Mobile,
		Email:    request.Email,
		Password: request.Password,
		Ip:       api.GetClientIP(ctx),
	})
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
