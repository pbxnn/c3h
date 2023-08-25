package data

import (
	"c3h/api/platform"
	"c3h/internal/biz"
	"c3h/internal/data/dao"
	"context"
	"math/rand"
	"strconv"

	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.CollectRepo = (*collectorRepo)(nil)

type collectorRepo struct {
	logger *log.Helper
	cli    platform.PlatformHTTPClient
	db     *gorm.DB
}

func NewCollectorRepo(cli platform.PlatformHTTPClient, db *gorm.DB, logger log.Logger) biz.CollectRepo {
	return &collectorRepo{
		logger: log.NewHelper(log.With(logger, "module", "collector-repo")),
		cli:    cli,
		db:     db,
	}
}

func (rp *collectorRepo) CollectData(ctx context.Context) ([]*dao.DataInfo, error) {

	//resp, err :=rp.cli.CollectData(ctx, nil)
	//if err != nil {
	//	return nil, err
	//}

	resp, err := rp.mockCollectData(ctx)
	if err != nil {
		rp.logger.Warnf("mock collect data err:%s", err.Error())
		return nil, err
	}
	return resp, nil
}

func (rp *collectorRepo) mockCollectData(ctx context.Context) ([]*dao.DataInfo, error) {
	var res []*dao.DataInfo
	err := rp.db.Model(res).Find(&res).Error
	if err != nil {
		return nil, err
	}

	for idx, item := range res {
		if item.Type != 2 {
			continue
		}

		res[idx].RealValue = rp.getRandomValue(item.SetValue, 0.5)
		if rp.getControlStatus(item.Key, res) {
			res[idx].CalcValue = rp.getRandomValue(item.SetValue, 0.3)
		}
	}

	return res, nil
}

func (rp *collectorRepo) getControlStatus(key string, list []*dao.DataInfo) bool {
	status := false
	for _, item := range list {
		if item.Key == key+"_control_switch" && item.RealValue == 1 {
			status = true
		}
	}

	return status
}

func (rp *collectorRepo) getRandomValue(setValue, scope float64) float64 {

	min := setValue * (1 - scope)
	max := setValue * (1 + scope)

	res, _ := strconv.ParseFloat(strconv.FormatFloat(rand.Float64()*(max-min)+min, 'f', 4, 64), 64)
	return res
}
