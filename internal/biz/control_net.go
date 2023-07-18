package biz

import "github.com/go-kratos/kratos/v2/log"

type ControlNetRepo interface {
}

type ControlNetUsecase struct {
	cnr    ControlNetRepo
	logger *log.Helper
}

func NewControlNetUsecase(cnr ControlNetRepo, logger log.Logger) *ControlNetUsecase {
	return &ControlNetUsecase{
		cnr:    cnr,
		logger: log.NewHelper(log.With(logger, "module", "control-net-usecase")),
	}
}
