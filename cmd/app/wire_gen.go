// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"app/internal/conf"
	"app/internal/module/greeter/biz"
	"app/internal/module/greeter/data"
	"app/internal/module/greeter/service"
	"app/internal/server"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(s *conf.Server, d *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	greeterRepo := data.NewGreeterRepo(logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	serverService := &server.Service{
		Greeter: greeterService,
	}
	grpcServer := server.NewGRPCServer(s, logger, serverService)
	httpServer := server.NewHTTPServer(s, logger, serverService)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
	}, nil
}
