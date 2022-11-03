package resource

import proto "user/api/qvbilam/user/v1"

type LoginResource struct {
	ID       int64  `json:"id"`
	Code     int64  `json:"code"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
}

func (r *LoginResource) Resource(response *proto.AccountResponse, token string) *LoginResource {
	r.ID = response.User.Id
	r.Code = response.User.Code
	r.Nickname = response.User.Nickname
	r.Avatar = response.User.Avatar
	r.Token = token
	return r
}
