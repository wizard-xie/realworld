package middleware

import (
	"context"

	km "github.com/go-kratos/kratos/v2/middleware"
)

const tokenKey = "Authorization"

func JWTAuthMiddleware() km.Middleware {
	return func(handler km.Handler) km.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			// if tr, ok := transport.FromServerContext(ctx); ok {
			// 	//token := tr.RequestHeader().Get(tokenKey)

			// }

			return handler(ctx, req)
		}
	}
}

func lint() {
	i := 8
	if i == 0 {

	} else if i == 2 {

	} else if i == 3 {

	}
}
