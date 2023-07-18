package data

import (
	"c3h/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.ControlNetRepo = (*controlNetRepo)(nil)

type controlNetRepo struct {
	logger *log.Helper
}

func NewControlNetRepo(logger log.Logger) biz.ControlNetRepo {
	return &controlNetRepo{
		logger: log.NewHelper(log.With(logger, "module", "new-control-net-repo")),
	}
}
