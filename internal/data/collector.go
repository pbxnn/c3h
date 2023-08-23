package data

import (
	"c3h/internal/biz"
	"c3h/internal/data/dao"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.CollectRepo = (*collectorRepo)(nil)

type collectorRepo struct {
	logger *log.Helper
	db     *gorm.DB
}

func NewCollectorRepo(logger log.Logger) biz.CollectRepo {
	return &collectorRepo{
		logger: log.NewHelper(log.With(logger, "module", "collector-repo")),
	}
}

func (rp *collectorRepo) CollectData(ctx context.Context) ([]*dao.DataInfo, error) {

	return []*dao.DataInfo{
		{},
	}, nil
}
