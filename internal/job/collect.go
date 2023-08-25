package job

import (
	"c3h/api/platform"
	"c3h/internal/biz"
	"c3h/internal/conf"
	"c3h/internal/data/dao"
	"c3h/internal/service"
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

	uc  *biz.CollectUsecase
	alu *biz.AuditLogUsecase
	s   *service.R401SService
}

func NewCollectJob(bc *conf.Bootstrap, logger log.Logger, uc *biz.CollectUsecase, alu *biz.AuditLogUsecase, s *service.R401SService) *CollectJob {
	sc, ok := bc.Job.Schedule["collect"]
	if !ok {
		sc = DefaultCollectSchedule
	}

	return &CollectJob{
		logger:   log.NewHelper(log.With(logger, "module", "collect-job")),
		schedule: sc,
		uc:       uc,
		alu:      alu,
		s:        s,
	}
}

func (cj *CollectJob) Run() {
	var err error
	var list []*dao.DataInfo
	var op = platform.OperationPlatformCollectData
	ctx := cj.initCtx(context.Background())

	defer func() {
		cj.alu.AddCronRecord(ctx, op, list, err)
	}()

	list, err = cj.uc.Collect(ctx)
	if err != nil {
		cj.logger.Warn("collect job run err:%s", err.Error())
	}

	cj.logger.Debugf("cron finished")

	//biz.CollectedChan <- ctx
	cj.s.SendMessage(ctx)
}

func (cj *CollectJob) initCtx(ctx context.Context) context.Context {
	spanName := CollectJobName + "_" + time.Now().Format("20060102150405")
	tracer := otel.Tracer(spanName)
	c, _ := tracer.Start(ctx, spanName)
	return c
}
