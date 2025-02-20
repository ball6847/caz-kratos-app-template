package greeter

import (
	"app/internal/greetersvc/biz"
	"app/internal/greetersvc/data"
	"app/internal/greetersvc/service"

	"github.com/google/wire"
)

// ProviderSet provides all dependencies for greeter module
var ProviderSet = wire.NewSet(
	data.NewGreeterRepo,
	service.NewGreeterService,
	biz.NewGreeterUsecase,
)
