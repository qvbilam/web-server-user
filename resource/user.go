package resource

import proto "user/api/qvbilam/user/v1"

type UsersResource struct {
	Total int64           `json:"total"`
	List  []*UserResource `json:"list"`
}

type UserResource struct {
	ID       int64  `json:"id"`
	Code     int64  `json:"code"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Gender   string `json:"gender"`
}

func (r *UsersResource) Resource(p *proto.UsersResponse) *UsersResource {
	res := UsersResource{}
	res.Total = p.Total
	var users []*UserResource
	if res.Total > 0 {
		for _, item := range p.Users {
			user := UserResource{}
			users = append(users, user.Resource(item))
		}

		res.List = users
	}
	return &res
}

func (r *UserResource) Resource(p *proto.UserResponse) *UserResource {
	res := UserResource{
		ID:       p.Id,
		Code:     p.Code,
		Nickname: p.Nickname,
		Avatar:   p.Avatar,
		Gender:   p.Gender,
	}

	return &res
}