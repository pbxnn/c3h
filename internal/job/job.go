package job

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/robfig/cron/v3"
)

var JobProviderSet = wire.NewSet(NewCron, NewCronLogger, NewCollectJob)

func NewCron(logger *CronLogger, cj *CollectJob) *cron.Cron {
	c := cron.New(
		cron.WithSeconds(),
		//cron.WithLogger(logger),
	)

	_, err := c.AddJob(cj.schedule, cj)
	if err != nil {
		logger.Error(err, "add collect job failed")
	}
	return c
}

type CronLogger struct {
	logger log.Logger
}

func NewCronLogger(logger log.Logger) *CronLogger {
	return &CronLogger{
		logger: logger,
	}
}

func (cl *CronLogger) Info(msg string, keysAndValues ...interface{}) {
	log.With(cl.logger, "msg", msg).Log(log.LevelInfo, keysAndValues)
}

func (cl *CronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	log.With(cl.logger, "msg", msg, "err", err).Log(log.LevelError, keysAndValues)
}
