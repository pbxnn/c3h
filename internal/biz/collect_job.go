package biz

import (
	"c3h/internal/data/dao"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type CollectRepo interface {
	CollectData(ctx context.Context) ([]*dao.DataInfo, error)
}

type CollectUsecase struct {
	logger *log.Helper
	cr     CollectRepo
	mr     ModuleRelationRepo
	dr     DataInfoRepo
}

func NewCollectUsecase(logger log.Logger, cr CollectRepo, mr ModuleRelationRepo, dr DataInfoRepo) *CollectUsecase {
	return &CollectUsecase{
		logger: log.NewHelper(log.With(logger, "module", "collect-usecase")),
		cr:     cr,
		mr:     mr,
		dr:     dr,
	}
}

func (uc *CollectUsecase) Run(ctx context.Context) error {
	list, err := uc.cr.CollectData(ctx)
	if err != nil {
		uc.logger.Warnf("collect data err:%s", err.Error())
		return err
	}

	err = uc.dr.BatchUpdate(ctx, list)
	if err != nil {
		return err
	}

	uc.mr.RefreshCache(ctx)
	return nil
}
