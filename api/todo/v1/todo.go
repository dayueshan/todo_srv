package v1

import (
	"todo-server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type Status int

const (
	StatusOK       Status = 0
	StatusDisabled Status = 1
)

type CreateReq struct {
	g.Meta  `path:"/to_do" method:"post" tags:"Todo" summary:"Create todo"`
	UserId  int64   `json:"user_id"`
	Content *string `v:"required|length:1,1024" dc:"todo content"`
}
type CreateRes struct {
	Id int64 `json:"id" dc:"todo id"`
}

type UpdateReq struct {
	g.Meta  `path:"/to_do/{id}" method:"put" tags:"Todo" summary:"Update todo"`
	Id      int64   `v:"required" dc:"todo id"`
	Content *string `v:"length:1,1024" dc:"todo Content"`
	Status  *Status `v:"in:0,1" dc:"todo status"`
}
type UpdateRes struct{}

type DeleteReq struct {
	g.Meta `path:"/to_do/{id}" method:"delete" tags:"Todo" summary:"Delete todo"`
	Id     int64 `v:"required" dc:"todo id"`
}
type DeleteRes struct{}

type GetOneReq struct {
	g.Meta `path:"/to_do/{id}" method:"get" tags:"Todo" summary:"Get one todo"`
	Id     int64 `v:"required" dc:"todo id"`
}
type GetOneRes struct {
	*entity.Todo `dc:"todo"`
}

type GetListReq struct {
	g.Meta `path:"/to_do" method:"get" tags:"Todo" summary:"Get todos"`
	UserId int64   `v:"required" dc:"todo user id"`
	Status *Status `v:"in:0,1" dc:"todo status"`
}
type GetListRes struct {
	List []*entity.Todo `json:"list" dc:"todo list"`
}
