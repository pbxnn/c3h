package data

import (
	"c3h/internal/data/dao"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type DataInfoRepo struct {
	logger *log.Helper
	db     *gorm.DB
}

func NewDataInfoRepo(db *gorm.DB, logger log.Logger) *DataInfoRepo {

	return &DataInfoRepo{
		logger: log.NewHelper(log.With(logger, "module", "data-info-repo")),
		db:     db,
	}
}

func (rp *DataInfoRepo) GetByKeys(ctx context.Context, keys []string) ([]*dao.DataInfo, error) {
	var err error
	var res []*dao.DataInfo
	err = rp.db.WithContext(ctx).Table(dao.DataInfoTblName).Where("key", keys).Order("id ASC").Find(&res).Error
	return res, err
}
