package biz

import "github.com/go-kratos/kratos/v2/log"

type ProductNetRepo interface {
}

type ProductNetUsecase struct {
	pnr    ProductNetRepo
	logger *log.Helper
}

func NewProductNetUsecase(pnr ProductNetRepo, logger log.Logger) *ProductNetUsecase {
	return &ProductNetUsecase{
		pnr:    pnr,
		logger: log.NewHelper(log.With(logger, "module", "product-net-usecase")),
	}
}
