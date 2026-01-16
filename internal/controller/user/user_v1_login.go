package user

import (
	"context"
	"errors"
	"time"
	"todo-server/internal/dao"
	"todo-server/internal/model/entity"

	"todo-server/api/user/v1"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("abc")

type Claims struct {
	UserID uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	var u *entity.User
	err = dao.User.Ctx(ctx).Where("provider", req.Provider).Where("provider_user_id", req.ProviderUserId).Limit(1).Scan(&u)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, errors.New("未找到此渠道id的用户")
	}

	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		UserID: u.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var token string
	token, err = t.SignedString(jwtKey)

	res = &v1.LoginRes{
		Token: token,
	}

	return
}
