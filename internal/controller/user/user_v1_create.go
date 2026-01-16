package user

import (
	"context"
	"errors"
	"todo-server/api/user/v1"
	"todo-server/internal/dao"
	"todo-server/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	u, e := dao.User.Ctx(ctx).Where(g.Map{"provider": req.Provider, "provider_user_id": req.ProviderUserId}).One()

	if e != nil {
		return nil, e
	}

	if u != nil {
		return nil, errors.New("此渠道id的用户已经注册")
	}

	insertId, err := dao.User.Ctx(ctx).Data(do.User{
		Provider:       req.Provider,
		ProviderUserId: req.ProviderUserId,
		Name:           req.Name,
		Email:          req.Email,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	res = &v1.CreateRes{
		Id: insertId,
	}
	return
}
