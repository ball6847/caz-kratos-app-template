//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"<%= go_module_name %>/internal/conf"
	"<%= go_module_name %>/internal/greetersvc"
	"<%= go_module_name %>/internal/server"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(s *conf.Bootstrap, d *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		greetersvc.ProviderSet,
		newApp,
	))
}
