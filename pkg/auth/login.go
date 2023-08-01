package auth

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
)

func Login(opts ...Option) middleware.Middleware {

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if err := checkLogin(ctx); err != nil {
				return nil, err
			}

			r := getRole(ctx)

			ctx = context.WithValue(ctx, SecurityUserRoleContextKey, r)
			return handler(ctx, req)
		}
	}
}

func checkLogin(ctx context.Context) error {

	return nil
}

func getRole(ctx context.Context) string {
	return "admin"
}
