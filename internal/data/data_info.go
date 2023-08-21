package data

import (
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
}

func NewDataInfoRepo(db *gorm.DB, logger log.Logger) biz.DataInfoRepo {

	return &dataInfoRepo{
		logger: log.NewHelper(log.With(logger, "module", "data-info-repo")),
		db:     db,
	}
}

func (rp *dataInfoRepo) GetByKeys(ctx context.Context, keys []string) ([]*dao.DataInfo, error) {
	var err error
	var res []*dao.DataInfo
	err = rp.db.WithContext(ctx).Table(dao.DataInfoTblName).Where("key", keys).Order("id ASC").Find(&res).Error
	return res, err
}

func (rp *dataInfoRepo) UpdateByKey(ctx context.Context, key string, up map[string]interface{}) (*dao.DataInfo, error) {
	var res dao.DataInfo
	err := rp.db.WithContext(ctx).Model(res).Where("key", key).Updates(up).First(&res).Error
	return &res, err
}
