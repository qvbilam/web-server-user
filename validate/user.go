package validate

type CreateValidate struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Nickname string `form:"nickname" json:"nickname" binding:"omitempty,min=1,max=10"`
	Password string `form:"password" json:"password" binding:"omitempty,min=6,max=10"`
	Avatar   string `form:"avatar" json:"avatar" binding:"omitempty,url"`
	Gender   string `form:"gender" json:"gender" binding:"omitempty,oneof=female male"`
}

type LoginValidate struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Type     int64  `form:"type" json:"type" binding:"required,oneof=1 2"` // 1.验证码， 2密码
	Password string `form:"password" json:"password" binding:"omitempty,min=6,max=10"`
	Code     string `form:"code" json:"code" binding:"omitempty,min=6,max=10"`
}

type SearchValidate struct {
	Keyword string `form:"keyword" json:"keyword" binding:"omitempty"`
	Sort    string `form:"sort" json:"sort" binding:"omitempty"`
	Page    int64  `form:"page" json:"page" binding:"omitempty,min=1"`
	PerPage int64  `form:"per_page" json:"per_page" binding:"omitempty,min=1,max=999"`
}
