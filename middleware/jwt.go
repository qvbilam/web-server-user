package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
	AuthJWT "user/auth/jwt"
	"user/global"
)

// Auth 验证jwt
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "请登陆",
			})
			// 在中间件中结束
			ctx.Abort()
			return
		}

		// Bearer TokenValue 获取空格后面部分 TokenValue
		token = strings.Split(token, " ")[1]

		j := NewJWT()
		user, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"msg": "登陆过期",
				})
				// 在中间件中结束
				ctx.Abort()
				return
			}

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未登陆",
				"err": err.Error(),
			})
			// 在中间件中结束
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Set("userId", user.ID)
		// 继续执行
		ctx.Next()
	}
}

type JWT struct {
	SingingKey []byte
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that is not even a token")
	TokenInvalid     = errors.New("could not handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.ServerConfig.JWTConfig.SigningKey),
	}
}

// CreateToken 创建token CustomClaims
func (j *JWT) CreateToken(claims AuthJWT.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SingingKey)
}

// ParseToken 解析token
func (j *JWT) ParseToken(tokenString string) (*AuthJWT.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthJWT.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SingingKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			return nil, TokenMalformed
		} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
			return nil, TokenExpired
		} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
			return nil, TokenNotValidYet
		} else {
			return nil, TokenInvalid
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*AuthJWT.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}
	return nil, TokenInvalid
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &AuthJWT.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SingingKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*AuthJWT.CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(time.Duration(global.ServerConfig.JWTConfig.Expire)).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
