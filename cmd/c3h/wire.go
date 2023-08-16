//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"c3h/internal/biz"
	"c3h/internal/conf"
	"c3h/internal/data"
	"c3h/internal/job"
	"c3h/internal/server"
	"c3h/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, *conf.Data, log.Logger) (*C3h, func(), error) {
	panic(
		wire.Build(
			server.SrvProviderSet,
			data.DProviderSet,
			biz.BProviderSet,
			service.SvcProviderSet,
			job.JobProviderSet,
			newApp,
		),
	)
}
