package global

import (
	ut "github.com/go-playground/universal-translator"
	proto "user/api/qvbilam/user/v1"
	"user/config"
)

var (
	Trans        ut.Translator // 表单验证
	ServerConfig *config.ServerConfig

	UserServerClient proto.UserClient
)
