package greeter

import (
	"<%= go_module_name %>/internal/greetersvc/biz"
	"<%= go_module_name %>/internal/greetersvc/data"
	"<%= go_module_name %>/internal/greetersvc/service"

	"github.com/google/wire"
)

// ProviderSet provides all dependencies for greeter module
var ProviderSet = wire.NewSet(
	data.NewGreeterRepo,
	service.NewGreeterService,
	biz.NewGreeterUsecase,
)
