package user

import (
	"context"
	"todo-server/internal/dao"

	"todo-server/api/user/v1"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	res = &v1.GetOneRes{}
	err = dao.User.Ctx(ctx).WherePri(req.Id).Scan(&res.User)
	return
}
