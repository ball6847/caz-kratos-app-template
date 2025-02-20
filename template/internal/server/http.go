package server

import (
	helloworldv1 "<%= go_module_name %>/api/helloworld/v1"
	"<%= go_module_name %>/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	swaggerUI "github.com/tx7do/kratos-swagger-ui"
)

// registerSwagger to http server, default to /api-docs
func registerSwagger(srv *http.Server, c *conf.Bootstrap) {
	if c.Server.Http.Swagger == nil || !c.Server.Http.Swagger.Enabled {
		return
	}
	print(c.Server.Http.Swagger)
	if c.Server.Http.Swagger.LocalFile == "" {
		c.Server.Http.Swagger.LocalFile = "openapi.yaml"
	}
	if c.Server.Http.Swagger.Title == "" {
		c.Server.Http.Swagger.Title = c.Name
	}
	if c.Server.Http.Swagger.BasePath == "" {
		c.Server.Http.Swagger.BasePath = "/api-docs"
	}
	swaggerUI.RegisterSwaggerUIServerWithOption(
		srv,
		swaggerUI.WithLocalFile(c.Server.Http.Swagger.LocalFile),
		swaggerUI.WithTitle(c.Server.Http.Swagger.Title),
		swaggerUI.WithBasePath(c.Server.Http.Swagger.BasePath),
	)
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, logger log.Logger, svc *Service) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	registerSwagger(srv, c)
	helloworldv1.RegisterGreeterHTTPServer(srv, svc.Greeter)
	return srv
}
