package server

import (
	"c3h/api/control_net"
	"c3h/api/product_net"
	"c3h/internal/biz"
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
	"go.opentelemetry.io/otel"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(bc *conf.Bootstrap, alu *biz.AuditLogUsecase, cn *service.ControlNetService, pn *service.ProductNetService, logger log.Logger) *http.Server {

	initTracer()

	var opts = []http.ServerOption{
		NewMiddleware(alu, logger),
		http.Filter(handlers.CORS(
			handlers.AllowCredentials(),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Accept", "Refer", "Sec-Ch-Ua", "Sec-Ch-Ua-Mobile", "Sec-Ch-Ua-Platform", "User-Agent"}),
			handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "HEAD", "DELETE"}),
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
func NewMiddleware(alu *biz.AuditLogUsecase, logger log.Logger) http.ServerOption {
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
		AuditLog(alu),
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

// 设置全局trace
func initTracer() {
	//// 创建 Jaeger exporter
	//exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	//if err != nil {
	//	return err
	//}
	tp := tracesdk.NewTracerProvider(
	//// 将基于父span的采样率设置为100%
	//tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
	//// 始终确保在生产中批量处理
	//tracesdk.WithBatcher(exp),
	//// 在资源中记录有关此应用程序的信息
	//tracesdk.WithResource(resource.NewSchemaless(
	//	semconv.ServiceNameKey.String("kratos-trace"),
	//	attribute.String("exporter", "jaeger"),
	//	attribute.Float64("float", 312.23),
	//)),
	)

	otel.SetTracerProvider(tp)
}
