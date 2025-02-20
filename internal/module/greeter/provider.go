package greeter

import (
	"app/internal/module/greeter/biz"
	"app/internal/module/greeter/data"
	"app/internal/module/greeter/service"

	"github.com/google/wire"
)

// ProviderSet provides all dependencies for greeter module
var ProviderSet = wire.NewSet(
	data.NewGreeterRepo,
	service.NewGreeterService,
	biz.NewGreeterUsecase,
)
