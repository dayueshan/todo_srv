package v1

import (
	"todo-server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta         `path:"/user" method:"post" tags:"User" summary:"Create user"`
	Provider       *string `v:"required|length:1,32" json:"provider"`
	ProviderUserId *string `v:"required|length:1,128" json:"provider user id"`
	Name           *string `v:"required|length:1,64" dc:"user name"`
	Email          *string `v:"required|length:1,64" dc:"user email"`
}
type CreateRes struct {
	Id int64 `json:"id" dc:"user id"`
}

type GetOneReq struct {
	g.Meta `path:"/user/{id}" method:"get" tags:"User" summary:"Get one user"`
	Id     int64 `v:"required" dc:"user id"`
}
type GetOneRes struct {
	*entity.User `dc:"user"`
}

type LoginReq struct {
	g.Meta         `path:"/user/login" method:"post" tags:"User" summary:"user login"`
	Provider       *string `v:"required" json:"provider"`
	ProviderUserId *string `v:"required" json:"provider user id"`
}
type LoginRes struct {
	Token string `json:"token" dc:"user token"`
}
