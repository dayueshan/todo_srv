// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"todo-server/api/user/v1"
)

type IUserV1 interface {
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error)
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
}
