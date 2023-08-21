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
	mrp    biz.ModuleRelationRepo
	drp    biz.DataInfoRepo
}

func NewR401SRepo(db *gorm.DB, mrp biz.ModuleRelationRepo, drp biz.DataInfoRepo, logger log.Logger) biz.R401SRepo {
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
		log.Warnf("get data info err:%s, keys:%v", err.Error(), dataKeyList)
		return nil, err
	}

	return dataList, nil
}

func (rp *r401sRepo) UpdateSwitchStatus(ctx context.Context, key string, status float64) (*dao.DataInfo, error) {

	up := map[string]interface{}{
		"real_value": status,
		"set_value":  status,
		"calc_value": status,
	}
	return rp.drp.UpdateByKey(ctx, key, up)
}

func (rp *r401sRepo) ResetVar(ctx context.Context, key string) (*dao.DataInfo, error) {
	//keyList := []string{key}
	//dataList, err := rp.drp.GetByKeys(ctx, keyList)
	//if err != nil {
	//	log.Warnf("get data info err:%s, keys:%v", err.Error(), keyList)
	//	return nil, err
	//}
	//
	//if len(dataList) == 0 {
	//	msg := fmt.Sprintf("data info not exist, keys:%v", keyList)
	//	log.Warnf(msg)
	//	return nil, errors.New(msg)
	//}

	up := map[string]interface{}{
		//"calc_value": dataList[0].SetValue,
		"calc_value": gorm.Expr("`set_value`"),
	}

	res, err := rp.drp.UpdateByKey(ctx, key, up)
	if err != nil {
		log.Warnf("reset data info err:%s, keys:%v", err.Error(), key)
		return nil, err
	}

	return res, nil
}

func (rp *r401sRepo) UpdateVarSetValueByKey(ctx context.Context, key string, setValue float64) (*dao.DataInfo, error) {
	up := map[string]interface{}{
		"set_value": setValue,
	}

	return rp.drp.UpdateByKey(ctx, key, up)
}
