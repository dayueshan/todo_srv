package todo

import (
	"context"
	"todo-server/internal/dao"

	"todo-server/api/todo/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	_, err = dao.Todo.Ctx(ctx).WherePri(req.Id).Delete()
	return
}
