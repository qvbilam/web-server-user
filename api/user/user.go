package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
	"user/api"
	pageProto "user/api/qvbilam/page/v1"
	proto "user/api/qvbilam/user/v1"
	"user/global"
	"user/resource"
	"user/validate"
)

func Update(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	request := validate.UpdateValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	client := global.UserServerClient

	_, err := client.Update(context.Background(), &proto.UpdateRequest{
		Id:       userID,
		Nickname: request.Nickname,
		Gender:   request.Gender,
		Avatar:   request.Avatar,
	})

	if err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}

func Detail(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, _ := strconv.Atoi(id)

	client := global.UserServerClient

	entity, err := client.Detail(context.Background(), &proto.GetUserRequest{
		Id: int64(userId),
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	res := resource.UserResource{}
	api.SuccessNotMessage(ctx, res.Resource(entity))
}

func Search(ctx *gin.Context) {
	request := validate.SearchValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	if request.Page <= 0 {
		request.Page = 1
	}
	if request.PerPage == 0 {
		request.PerPage = 10
	}

	client := global.UserServerClient
	res, err := client.Search(context.Background(), &proto.SearchRequest{
		Keyword: request.Keyword,
		Sort:    request.Sort,
		Page: &pageProto.PageRequest{
			Page:    request.Page,
			PerPage: request.PerPage,
		},
	})

	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	var users []interface{}
	for _, user := range res.Users {
		users = append(users, proto.UserResponse{
			Id:       user.Id,
			Code:     user.Code,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
			Gender:   user.Gender,
			Level:    nil,
		})
	}

	usersResource := resource.UsersResource{}
	api.SuccessNotMessage(ctx, usersResource.Resource(res))
}

func Delete(ctx *gin.Context) {
}
