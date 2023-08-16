package server

import (
	"c3h/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
)

func AuditLog(alu *biz.AuditLogUsecase) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {

			reply, err = handler(ctx, req)

			//异步记录审计日志
			go alu.AddRequestRecord(ctx, req, reply, err)
			return
		}
	}
}
