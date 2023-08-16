package job

import (
	"c3h/api/platform"
	"c3h/internal/biz"
	"c3h/internal/conf"
	"c3h/internal/data/dao"
	"context"
	"time"

	"go.opentelemetry.io/otel"

	"github.com/go-kratos/kratos/v2/log"
)

const DefaultCollectSchedule = "*/10 * * * * *"

const CollectJobName = "c3h_collector"

type CollectJob struct {
	logger   *log.Helper
	schedule string
	cr       biz.CollectRepo
	alu      *biz.AuditLogUsecase
}

func NewCollectJob(bc *conf.Bootstrap, logger log.Logger, cr biz.CollectRepo, alu *biz.AuditLogUsecase) *CollectJob {
	sc, ok := bc.Job.Schedule["collect"]
	if !ok {
		sc = DefaultCollectSchedule
	}

	return &CollectJob{
		logger:   log.NewHelper(log.With(logger, "module", "collect-service")),
		schedule: sc,
		cr:       cr,
		alu:      alu,
	}
}

func (cj *CollectJob) Run() {
	var err error
	var list []*dao.DataInfo
	var op = platform.OperationProductNetCollectData
	ctx := cj.initCtx(context.Background())

	defer func() {
		cj.alu.AddCronRecord(ctx, op, list, err)
	}()

	list, err = cj.cr.CollectData(ctx)
	if err != nil {
		cj.logger.Warn()
		return
	}

	err = cj.cr.UpdateDB(ctx, list)
	if err != nil {
		return
	}

	err = cj.cr.UpdateCache(ctx, list)
	if err != nil {
		return
	}
}

func (cj *CollectJob) initCtx(ctx context.Context) context.Context {
	spanName := CollectJobName + "_" + time.Now().Format("20060102150405")
	tracer := otel.Tracer(spanName)
	c, _ := tracer.Start(ctx, spanName)
	return c
}
