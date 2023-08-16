package data

import (
	"context"

	"gorm.io/gorm"

	"c3h/internal/biz"
	"c3h/internal/data/dao"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.R401SRepo = (*r401sRepo)(nil)

type r401sRepo struct {
	logger *log.Helper
	db     *gorm.DB
	mrp    *ModuleRelationRepo
	drp    *DataInfoRepo
}

func NewR401SRepo(db *gorm.DB, mrp *ModuleRelationRepo, drp *DataInfoRepo, logger log.Logger) biz.R401SRepo {
	return &r401sRepo{
		logger: log.NewHelper(log.With(logger, "module", "r401s-repo")),
		db:     db,
		mrp:    mrp,
		drp:    drp,
	}
}

func (rp *r401sRepo) GetVarsByModule(ctx context.Context, moduleKey string) ([]*dao.DataInfo, error) {
	var err error
	var moduleRelationList []*dao.ModuleDataMap
	moduleRelationList, err = rp.mrp.GetByModuleKey(ctx, moduleKey)
	if err != nil {
		log.Warnf("get module relation err:%s, key:%s", err.Error(), moduleKey)
		return nil, err
	}

	var dataList []*dao.DataInfo
	var dataKeyList []string
	for _, item := range moduleRelationList {
		dataKeyList = append(dataKeyList, item.DataKey)
	}
	dataList, err = rp.drp.GetByKeys(ctx, dataKeyList)
	if err != nil {
		log.Warnf("get module relation err:%s, keys:%v", err.Error(), dataKeyList)
		return nil, err
	}

	return dataList, nil
}

func (rp *r401sRepo) UpdateSwitchStatus(ctx context.Context, key string, status float64) (*dao.DataInfo, error) {

	return nil, nil
}

func (rp *r401sRepo) ResetVar(ctx context.Context, key string) (*dao.DataInfo, error) {

	return nil, nil
}
