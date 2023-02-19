package account

import (
	"context"
	"github.com/gin-gonic/gin"
	"user/api"
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
	if _, ok := passwordMethod[request.Method]; ok {
		entity, err = global.AccountServerClient.LoginPassword(context.Background(), &proto.LoginPasswordRequest{
			Method:   request.Method,
			Username: request.Username,
			Mobile:   request.Mobile,
			Email:    request.Email,
			Password: request.Password,
			Ip:       api.GetClientIP(ctx),
		})
		if err != nil {
			api.HandleGrpcErrorToHttp(ctx, err)
			return
		}
	} else {
		api.ErrorUnprocessableEntity(ctx, gin.H{"method": "Unprocessable method"})
		return
	}

	token := entity.Token

	r := resource.LoginResource{}
	api.SuccessNotMessage(ctx, r.Resource(entity, token))
}
