package server

import (
	"c3h/api/control_net"
	"c3h/api/product_net"
	"c3h/internal/conf"
	"c3h/internal/service"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/handlers"
	"gorm.io/gorm"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(bc *conf.Bootstrap, db *gorm.DB, cn *service.ControlNetService, pn *service.ProductNetService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		NewMiddleware(bc.Auth, db, logger),
		http.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST"}),
		)),
		//http.ResponseEncoder(encoder.ResponseEncoder),
		//http.ErrorEncoder(encoder.ErrorEncoder),
	}

	c := bc.Server
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	product_net.RegisterProductNetHTTPServer(srv, pn)
	control_net.RegisterControlNetHTTPServer(srv, cn)

	openAPIHandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIHandler)

	return srv
}

// NewMiddleware 创建中间件
func NewMiddleware(ac *conf.Auth, db *gorm.DB, logger log.Logger) http.ServerOption {
	//m, err := model.NewModelFromFile(ac.Model)
	//if err != nil {
	//	panic(err)
	//}
	//
	//a, err := gormadapter.NewAdapterByDB(db)
	//if err != nil {
	//	panic(err)
	//}

	return http.Middleware(
		recovery.Recovery(),
		tracing.Server(),
		logging.Server(logger),
		//selector.Server(
		//	auth.Login(),
		//	auth.Casbin(
		//		auth.WithCasbinModel(m),
		//		auth.WithCasbinPolicy(a),
		//		auth.WithSecurityUserCreator(auth.NewSecurityUser),
		//	),
		//).Match(NewWhiteListMatcher(ac)).Build(),
	)
}

// NewWhiteListMatcher 创建白名单
func NewWhiteListMatcher(ac *conf.Auth) selector.MatchFunc {

	return func(ctx context.Context, operation string) bool {
		tr, ok := transport.FromServerContext(ctx)
		if !ok {
			return true
		}

		wh := tr.RequestHeader().Get(ac.WhiteHeader)
		if _, ok := ac.WhiteMap[wh]; ok {
			return false
		}

		return true
	}
}
