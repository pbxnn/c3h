package data

import (
	"c3h/pkg/cache"
	"context"

	"c3h/internal/biz"
	"c3h/internal/data/dao"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.DataInfoRepo = (*dataInfoRepo)(nil)

type dataInfoRepo struct {
	logger *log.Helper
	db     *gorm.DB
	cache  cache.Cache
}

func NewDataInfoRepo(db *gorm.DB, c cache.Cache, logger log.Logger) biz.DataInfoRepo {
	return &dataInfoRepo{
		logger: log.NewHelper(log.With(logger, "module", "data-info-repo")),
		db:     db,
		cache:  c,
	}
}

func (rp *dataInfoRepo) GetByKeys(ctx context.Context, keys []string) ([]*dao.DataInfo, error) {
	var err error
	var res []*dao.DataInfo

	if cacheData, ok := rp.getByKeysFromCache(ctx, keys); ok {
		return cacheData, nil
	}

	err = rp.db.WithContext(ctx).Table(dao.DataInfoTblName).Where("key", keys).Order("id ASC").Find(&res).Error

	rp.updateCacheByKeys(ctx, res)
	return res, err
}

func (rp *dataInfoRepo) UpdateByKey(ctx context.Context, key string, up map[string]interface{}) (*dao.DataInfo, error) {
	var res dao.DataInfo
	err := rp.db.WithContext(ctx).Model(res).Where("key", key).Updates(up).First(&res).Error

	if err != nil {
		rp.updateCacheByKeys(ctx, []*dao.DataInfo{&res})
	}
	return &res, err
}

func (rp *dataInfoRepo) getByKeysFromCache(ctx context.Context, keys []string) ([]*dao.DataInfo, bool) {
	var err error
	var hit = true
	var res []*dao.DataInfo
	for _, k := range keys {
		var item *dao.DataInfo
		err = rp.cache.Get(k, item)
		if err != nil || item == nil {
			hit = false
			break
		}
		res = append(res, item)
	}

	return res, hit
}

func (rp *dataInfoRepo) updateCacheByKeys(ctx context.Context, cacheData []*dao.DataInfo) {
	for _, item := range cacheData {
		rp.cache.Set(item.Key, item)
	}
}

func (rp *dataInfoRepo) BatchUpdate(ctx context.Context, list []*dao.DataInfo) error {
	for _, item := range list {
		up := map[string]interface{}{
			"real_value": item.RealValue,
			"set_value":  item.SetValue,
			"calc_value": item.CalcValue,
			"high_limit": item.HighLimit,
			"low_limit":  item.LowLimit,
		}
		_, err := rp.UpdateByKey(ctx, item.Key, up)
		if err != nil {
			return err
		}
	}

	return nil
}
