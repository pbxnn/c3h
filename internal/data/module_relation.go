package data

import (
	"c3h/internal/biz"
	"c3h/internal/data/dao"
	"c3h/pkg/cache"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.ModuleRelationRepo = (*moduleRelationRepo)(nil)

const ModuleMapCacheKey = "ModuleMapKey_"
const DataMapCacheKey = "DataMapKey_"

type moduleRelationRepo struct {
	logger *log.Helper
	db     *gorm.DB
	cache  cache.Cache
}

func NewModuleRelationRepo(db *gorm.DB, cache cache.Cache, logger log.Logger) biz.ModuleRelationRepo {
	return &moduleRelationRepo{
		logger: log.NewHelper(log.With(logger, "module", "module-relation-repo")),
		db:     db,
		cache:  cache,
	}
}

func (rp *moduleRelationRepo) RefreshCache(ctx context.Context) {
	var res []*dao.ModuleDataMap
	var err error

	err = rp.db.Model(res).Find(&res).Error
	if err != nil {
		rp.logger.Warn("refresh cache failed. get data from db err:%s")
		return
	}

	rp.setCache(ctx, res)
}

func (rp *moduleRelationRepo) GetByModuleKey(ctx context.Context, moduleKey string) ([]*dao.ModuleDataMap, error) {
	var res []*dao.ModuleDataMap
	var err error

	cacheData, err := rp.getByModuleKeyFromCache(ctx, moduleKey)
	if err == nil && cacheData != nil {
		return cacheData, nil
	}

	err = rp.db.WithContext(ctx).Model(res).Where("module_key", moduleKey).Find(&res).Error
	rp.setCache(ctx, res)

	return res, err
}

func (rp *moduleRelationRepo) getByModuleKeyFromCache(ctx context.Context, moduleKey string) ([]*dao.ModuleDataMap, error) {
	var res []*dao.ModuleDataMap
	err := rp.cache.Get(ModuleMapCacheKey+moduleKey, res)
	if err != nil || len(res) == 0 {
		return nil, fmt.Errorf("get module relation real from cache err:%v, key:%s", err, moduleKey)
	}
	return res, err
}

func (rp *moduleRelationRepo) GetByDataKeys(ctx context.Context, dataKeys []string) ([]*dao.ModuleDataMap, error) {
	var res []*dao.ModuleDataMap
	var err error

	cacheData, err := rp.getByDataKeysFromCache(ctx, dataKeys)
	if err == nil && cacheData != nil {
		return cacheData, nil
	}

	err = rp.db.WithContext(ctx).Model(res).Where("data_key", dataKeys).Find(&res).Error
	rp.setCache(ctx, res)

	return res, err
}

func (rp *moduleRelationRepo) getByDataKeysFromCache(ctx context.Context, dataKeys []string) ([]*dao.ModuleDataMap, error) {
	var res []*dao.ModuleDataMap
	for _, key := range dataKeys {
		var temp []*dao.ModuleDataMap
		err := rp.cache.Get(key, temp)
		if err != nil || len(res) == 0 {
			return nil, fmt.Errorf("get module relation real from cache err:%v, key:%s", err, key)
		}
		res = append(res, temp...)
	}

	return res, nil
}

func (rp *moduleRelationRepo) setCache(ctx context.Context, res []*dao.ModuleDataMap) {
	moduleKeyMap := map[string][]*dao.ModuleDataMap{}
	dataKeyMap := map[string][]*dao.ModuleDataMap{}

	for idx, item := range res {
		moduleKeyMap[item.ModuleKey] = append(moduleKeyMap[item.ModuleKey], res[idx])
		dataKeyMap[item.DataKey] = append(dataKeyMap[item.DataKey], res[idx])
	}

	for k, v := range moduleKeyMap {
		rp.cache.Set(ModuleMapCacheKey+k, v)
	}

	for k, v := range dataKeyMap {
		rp.cache.Set(DataMapCacheKey+k, v)
	}
}
