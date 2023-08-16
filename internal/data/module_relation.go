package data

import (
	"c3h/internal/data/dao"
	"c3h/pkg/cache"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type ModuleRelationRepo struct {
	logger *log.Helper
	db     *gorm.DB
	cache  cache.Cache
}

func NewModuleRelationRepo(db *gorm.DB, cache cache.Cache, logger log.Logger) *ModuleRelationRepo {
	return &ModuleRelationRepo{
		logger: log.NewHelper(log.With(logger, "module", "module-relation-repo")),
		db:     db,
		cache:  cache,
	}
}

func (mr *ModuleRelationRepo) UpdateModuleRelationCache(ctx context.Context) {

	var dataList []*dao.DataInfo
	if err := mr.db.Find(&dataList).WithContext(ctx).Error; err != nil {
		mr.logger.Warn()
		return
	}

	var moduleList []*dao.ModuleInfo
	if err := mr.db.Find(&moduleList).WithContext(ctx).Error; err != nil {
		mr.logger.Warn()
		return
	}

	var mdList []*dao.ModuleDataMap
	err := mr.db.Find(&mdList).WithContext(ctx).Error
	if err != nil {
		mr.logger.Warn()
		return
	}

	moduleMap := map[uint]*dao.ModuleInfo{}
	for idx, item := range moduleList {
		moduleMap[item.ID] = moduleList[idx]
	}

	dataMap := map[uint]*dao.DataInfo{}
	for idx, item := range dataList {
		dataMap[item.ID] = dataList[idx]
	}

	moduleRelation := map[string][]string{}
	for _, item := range mdList {
		var mKey string
		var dKey string
		if m, ok := moduleMap[item.ModuleID]; ok {
			mKey = m.Key
		}

		if d, ok := dataMap[item.DataID]; ok {
			dKey = d.Key
		}

		if len(dKey) == 0 || len(mKey) == 0 {
			continue
		}

		if _, ok := moduleRelation[mKey]; !ok {
			moduleRelation[mKey] = []string{}
		}
		moduleRelation[mKey] = append(moduleRelation[mKey], dKey)
	}

	for k, v := range moduleRelation {
		if err := mr.cache.Set(k, v); err != nil {
			mr.logger.Warn()
		}
	}
}