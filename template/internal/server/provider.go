package server

import (
	greetersvc "<%= go_module_name %>/internal/greetersvc/service"

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
