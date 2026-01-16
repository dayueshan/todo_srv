package middleware

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(r *ghttp.Request) {
	path := r.URL.Path

	// 白名单：登录接口不需要认证
	if path == "/user/login" {
		r.Middleware.Next()
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		r.Response.WriteStatusExit(401, "缺少 Authorization header")
		return
	}

	tokenStr := authHeader
	if gstr.HasPrefix(authHeader, "Bearer ") {
		tokenStr = authHeader[7:]
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte("abc"), nil
	})

	if err != nil || !token.Valid {
		r.Response.WriteStatusExit(401, "token不合法")
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if uid, exists := claims["user_id"]; exists {
			// 将 user_id 存入 context，供后续使用
			ctx := context.WithValue(r.Context(), "user_id", uint64(uid.(float64)))
			r.SetCtx(ctx)
			r.Middleware.Next()
			return
		}
	}

	r.Response.WriteStatusExit(401, "Token 缺少 user_id")
}
