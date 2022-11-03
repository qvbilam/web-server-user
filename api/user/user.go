package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"user/api"
	pageProto "user/api/qvbilam/page/v1"
	proto "user/api/qvbilam/user/v1"
	"user/global"
	"user/validate"
)

func Update(ctx *gin.Context) {
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

	//u, _ := ctx.Get("user")
	//model := u.(*AuthJWT.CustomClaims)
	//model.Id

	ctx.JSON(http.StatusOK, gin.H{
		"id":       entity.Id,
		"code":     entity.Code,
		"nickname": entity.Nickname,
		"avatar":   entity.Avatar,
	})
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

	ctx.JSON(http.StatusOK, gin.H{
		"total":    res.Total,
		"data":     users,
		"page":     request.Page,
		"per_page": request.PerPage,
	})
}

func Delete(ctx *gin.Context) {
}
