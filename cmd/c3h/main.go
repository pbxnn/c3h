package main

import (
	"flag"
	"os"

	"github.com/robfig/cron/v3"

	"c3h/internal/conf"
	myLog "c3h/pkg/log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

type C3h struct {
	Cron *cron.Cron
	App  *kratos.App
}

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, c *cron.Cron) *C3h {
	return &C3h{
		Cron: c,
		App: kratos.New(
			kratos.ID(id),
			kratos.Name(Name),
			kratos.Version(Version),
			kratos.Metadata(map[string]string{}),
			kratos.Logger(logger),
			kratos.Server(
				hs,
			),
		),
	}
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)

	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	Name = bc.Name
	Version = bc.Version

	logger := log.With(
		myLog.NewLogrusLogger(
			myLog.Level(bc.Log.Level),
			myLog.FileOutput(bc.Log.Filename, bc.Log.RotateTime, bc.Log.MaxAge),
		),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	c3h, cleanup, err := wireApp(&bc, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start cron job
	c3h.Cron.Start()

	// start and wait for stop signal
	if err := c3h.App.Run(); err != nil {
		panic(err)
	}
}
