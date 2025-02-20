//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"app/internal/conf"
	"app/internal/greetersvc"
	"app/internal/server"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(s *conf.Server, d *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		greeter.ProviderSet,
		newApp,
	))
}
