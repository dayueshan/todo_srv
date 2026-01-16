package todo

import (
	"context"
	"todo-server/internal/dao"
	"todo-server/internal/model/do"

	"todo-server/api/todo/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	_, err = dao.Todo.Ctx(ctx).Data(do.Todo{
		Content: req.Content,
		Status:  req.Status,
	}).WherePri(req.Id).Update()
	return
}
