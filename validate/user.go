package validate

type SearchValidate struct {
	Keyword string `form:"keyword" json:"keyword" binding:"omitempty"`
	Sort    string `form:"sort" json:"sort" binding:"omitempty"`
	Page    int64  `form:"page" json:"page" binding:"omitempty,min=1"`
	PerPage int64  `form:"per_page" json:"per_page" binding:"omitempty,min=1,max=999"`
}

type UpdateValidate struct {
	Nickname string `form:"nickname" json:"nickname" binding:"omitempty"`
	Gender   string `form:"gender" json:"gender" binding:"omitempty,oneof=female male"`
	Avatar   string `form:"avatar" json:"avatar" binding:"omitempty,url"`
}
