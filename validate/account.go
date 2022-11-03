package validate

type CreateValidate struct {
	Username string `form:"username" json:"username" binding:"omitempty,min=4,max=20"`
	Mobile   string `form:"mobile" json:"mobile" binding:"omitempty"`
	Email    string `form:"email" json:"email" binding:"omitempty"`
	Password string `form:"password" json:"password" binding:"omitempty,min=6,max=10"`
}

type LoginValidate struct {
	Username string `form:"username" json:"username" binding:"omitempty"`
	Email    string `form:"email" json:"email" binding:"omitempty"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	//Type     int64  `form:"type" json:"type" binding:"required,oneof=1 2"` // 1.验证码， 2密码
	Method   string `form:"method" json:"method" binding:"required"` // 登陆方式
	Password string `form:"password" json:"password" binding:"omitempty,min=6,max=10"`
	Code     string `form:"code" json:"code" binding:"omitempty,min=6,max=10"`
}
