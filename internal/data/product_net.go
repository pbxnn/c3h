package data

import (
	"c3h/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.ProductNetRepo = (*productNetRepo)(nil)

type productNetRepo struct {
	logger *log.Helper
}

func NewProductNetRepo(logger log.Logger) biz.ProductNetRepo {
	return &productNetRepo{
		logger: log.NewHelper(log.With(logger, "module", "new-product-repo")),
	}
}
