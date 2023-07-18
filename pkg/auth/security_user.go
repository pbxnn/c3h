package auth

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/go-kratos/kratos/v2/transport"
	"github.com/tx7do/kratos-casbin/authz"
)

type SecurityUser struct {
	Path   string
	Method string
	Role   string
	Domain string
}

func NewSecurityUser() authz.SecurityUser {
	return &SecurityUser{}
}

func (su *SecurityUser) ParseFromContext(ctx context.Context) error {
	if tr, ok := transport.FromServerContext(ctx); ok {
		su.Path = tr.Operation()
		su.Method = "*"
		if r := ctx.Value(SecurityUserRoleContextKey); r != nil {
			su.Role = r.(string)
		}
		if tr.Kind() == transport.KindHTTP {
			if info, ok := tr.(*http.Transport); ok {
				req := info.Request()
				su.Method = req.Method
			}
		}

	} else {
		return errors.New("parse server context error")
	}

	return nil
}

func (su *SecurityUser) GetSubject() string {
	return su.Role
}

func (su *SecurityUser) GetObject() string {
	return su.Path
}

func (su *SecurityUser) GetAction() string {
	return su.Method
}

func (su *SecurityUser) GetDomain() string {
	return su.Domain
}
