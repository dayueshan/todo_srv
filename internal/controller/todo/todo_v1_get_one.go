package todo

import (
	"context"
	"todo-server/internal/dao"

	"todo-server/api/todo/v1"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	res = &v1.GetOneRes{}
	err = dao.Todo.Ctx(ctx).WherePri(req.Id).Scan(&res.Todo)
	return
}
