package server

import (
	helloworldv1 "app/api/helloworld/v1"
	"app/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	swaggerUI "github.com/tx7do/kratos-swagger-ui"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, svc *Service) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
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
	if c.Http.Swagger {
		swaggerUI.RegisterSwaggerUIServerWithOption(
			srv,
			swaggerUI.WithLocalFile("openapi.yaml"),
			swaggerUI.WithTitle("Hello API"),
			swaggerUI.WithBasePath("/api-docs"),
		)
	}
	helloworldv1.RegisterGreeterHTTPServer(srv, svc.Greeter)
	return srv
}
