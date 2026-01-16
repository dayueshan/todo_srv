package todo

import (
	"context"
	"todo-server/internal/dao"
	"todo-server/internal/model/do"

	"todo-server/api/todo/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	insertId, err := dao.Todo.Ctx(ctx).Data(do.Todo{
		UserId:  req.UserId,
		Content: req.Content,
		Status:  v1.StatusOK,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	res = &v1.CreateRes{
		Id: insertId,
	}
	return
}
