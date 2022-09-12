package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	proto "user/api/qvbilam/user/v1"
	"user/global"
	"user/validate"
)

func Create(ctx *gin.Context) {
	request := validate.Create{}
	if err := ctx.BindQuery(&request); err != nil {
		HandleValidateError(ctx, err)
		return
	}

	client := global.UserServerClient
	entity, err := client.Create(context.Background(), &proto.SignInRequest{
		Nickname: request.Nickname,
		Mobile:   request.Mobile,
		Password: request.Password,
		Gender:   request.Gender,
	})
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id": entity.Id,
	})
}

func Update(ctx *gin.Context) {

}

func Detail(ctx *gin.Context) {

}

func List(ctx *gin.Context) {

}

func CheckPassword(ctx *gin.Context) {

}
