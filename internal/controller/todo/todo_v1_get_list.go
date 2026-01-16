package todo

import (
	"context"
	"todo-server/internal/dao"
	"todo-server/internal/model/do"

	"todo-server/api/todo/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res = &v1.GetListRes{}
	err = dao.Todo.Ctx(ctx).Where(do.Todo{
		UserId: req.UserId,
		Status: req.Status,
	}).Scan(&res.List)
	return
}
