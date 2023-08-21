package biz

import (
	"c3h/internal/data/dao"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type R401SRepo interface {
	GetVarsByModule(ctx context.Context, moduleKey string) ([]*dao.DataInfo, error)
	UpdateSwitchStatus(ctx context.Context, key string, status float64) (*dao.DataInfo, error)
	UpdateVarSetValueByKey(ctx context.Context, key string, setValue float64) (*dao.DataInfo, error)
	ResetVar(ctx context.Context, key string) (*dao.DataInfo, error)
}

type R401SUsecase struct {
	rp     R401SRepo
	drp    DataInfoRepo
	logger *log.Helper
}

func NewR401SUsecase(rp R401SRepo, drp DataInfoRepo, logger log.Logger) *R401SUsecase {
	return &R401SUsecase{
		rp:     rp,
		drp:    drp,
		logger: log.NewHelper(log.With(logger, "module", "R401S-usecase")),
	}
}

func (uc *R401SUsecase) GetVarsByModule(ctx context.Context, moduleKey string) ([]*dao.DataInfo, error) {
	return uc.rp.GetVarsByModule(ctx, moduleKey)
}

//func (uc *R401SUsecase) UpdateVarSetValue(ctx context.Context, key string, setValue float64) (*dao.DataInfo, error) {
//	return uc.rp.UpdateVarSetValueByKey(ctx, key, setValue)
//}

func (uc *R401SUsecase) UpdateSwitchStatus(ctx context.Context, key string, status float64) (*dao.DataInfo, error) {
	return uc.rp.UpdateSwitchStatus(ctx, key, status)
}

func (uc *R401SUsecase) ResetVar(ctx context.Context, key string) (*dao.DataInfo, error) {
	return uc.rp.ResetVar(ctx, key)
}

func (uc *R401SUsecase) UpdateControlSwitchStatus(ctx context.Context, key, controlKey string, status float64) ([]*dao.DataInfo, error) {

	controlSwitchInfo, err := uc.rp.UpdateSwitchStatus(ctx, controlKey, status)
	if err != nil {
		return nil, err
	}

	dataInfo, err := uc.drp.GetByKeys(ctx, []string{key})
	if err != nil {
		return nil, err
	}

	dataInfo = append(dataInfo, controlSwitchInfo)
	return dataInfo, nil
}
