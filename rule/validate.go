package rule

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateMobile(f validator.FieldLevel) bool {
	mobile := f.Field().String() // 获取获取手机号
	// 自定义表达式匹配
	if ok, _ := regexp.MatchString(`/^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\\d{8}$/`, mobile); !ok {
		return true
	}

	return true
}
