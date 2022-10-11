package api

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	proto "user/api/qvbilam/user/v1"
	AuthJWT "user/auth/jwt"
	"user/global"
	"user/middleware"
	"user/validate"
)

func Register(ctx *gin.Context) {
	request := validate.CreateValidate{}
	if err := ctx.BindQuery(&request); err != nil {
		HandleValidateError(ctx, err)
		return
	}

	client := global.UserServerClient
	entity, err := client.Create(context.Background(), &proto.UpdateRequest{
		Nickname: request.Nickname,
		Mobile:   request.Mobile,
		Password: request.Password,
		Gender:   request.Gender,
	})
	if err != nil {
		return
	}

	token := generateUserToken(entity)

	ctx.JSON(http.StatusOK, gin.H{
		"id":       entity.Id,
		"code":     entity.Code,
		"nickname": entity.Nickname,
		"avatar":   entity.Avatar,
		"token":    token,
	})
}

func Login(ctx *gin.Context) {
	request := validate.LoginValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		HandleValidateError(ctx, err)
		return
	}

	client := global.UserServerClient
	entity, err := client.Login(context.Background(), &proto.LoginRequest{
		Mobile:   request.Mobile,
		Password: request.Password,
	})

	if err != nil {
		HandleGrpcErrorToHttp(ctx, err)
		return
	}

	token := generateUserToken(entity)
	ctx.JSON(http.StatusOK, gin.H{
		"id":       entity.Id,
		"code":     entity.Code,
		"nickname": entity.Nickname,
		"avatar":   entity.Avatar,
		"token":    token,
	})
}

func Logout(ctx *gin.Context) {

}

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
		HandleGrpcErrorToHttp(ctx, err)
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
		HandleValidateError(ctx, err)
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
		Page: &proto.PageRequest{
			Page:    request.Page,
			PerPage: request.PerPage,
		},
	})

	if err != nil {
		HandleGrpcErrorToHttp(ctx, err)
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

func generateUserToken(entity *proto.UserResponse) string {
	j := middleware.NewJWT()
	expire := time.Now().Unix() + global.ServerConfig.JWTConfig.Expire
	issuer := global.ServerConfig.JWTConfig.Issuer
	claims := AuthJWT.CustomClaims{
		ID:       entity.Id,
		Code:     entity.Code,
		Nickname: entity.Nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire,
			Issuer:    issuer,
			NotBefore: time.Now().Unix(),
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		return ""
	}

	return token
}
