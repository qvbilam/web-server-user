package validate

type SearchValidate struct {
	Keyword string `form:"keyword" json:"keyword" binding:"omitempty"`
	Sort    string `form:"sort" json:"sort" binding:"omitempty"`
	Page    int64  `form:"page" json:"page" binding:"omitempty,min=1"`
	PerPage int64  `form:"per_page" json:"per_page" binding:"omitempty,min=1,max=999"`
}
