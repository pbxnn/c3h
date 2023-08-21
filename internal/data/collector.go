package data

import (
	"c3h/internal/biz"
	"c3h/internal/data/dao"
	"c3h/pkg/cache"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.CollectRepo = (*collectorRepo)(nil)

type collectorRepo struct {
	logger *log.Helper
	db     *gorm.DB
	cache  cache.Cache
	mr     biz.ModuleRelationRepo
}

func NewCollectorRepo(mr biz.ModuleRelationRepo, db *gorm.DB, cache cache.Cache, logger log.Logger) biz.CollectRepo {
	return &collectorRepo{
		logger: log.NewHelper(log.With(logger, "module", "collector-repo")),
		db:     db,
		cache:  cache,
		mr:     mr,
	}
}

func (cr *collectorRepo) CollectData(ctx context.Context) ([]*dao.DataInfo, error) {

	return []*dao.DataInfo{
		{},
	}, nil
}

func (cr *collectorRepo) UpdateDB(ctx context.Context, list []*dao.DataInfo) error {

	for _, item := range list {
		err := cr.db.Updates(item).Where("key", item.Key).Error
		if err != nil {
			cr.logger.Warn()
		}
	}

	return nil
}

func (cr *collectorRepo) UpdateCache(ctx context.Context, list []*dao.DataInfo) error {
	cr.mr.UpdateModuleRelationCache(ctx)
	cr.UpdateDataCache(ctx, list)

	return nil
}

func (cr *collectorRepo) UpdateDataCache(ctx context.Context, list []*dao.DataInfo) {
	for idx, item := range list {
		cr.cache.Set(item.Key, list[idx])
	}
}
