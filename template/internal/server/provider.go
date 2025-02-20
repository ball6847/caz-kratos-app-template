package server

import (
	greetersvc "app/internal/greetersvc/service"

	"github.com/google/wire"
)

type Service struct {
	Greeter *greetersvc.GreeterService
}

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	wire.Struct(new(Service), "*"),
	NewGRPCServer,
	NewHTTPServer,
)
