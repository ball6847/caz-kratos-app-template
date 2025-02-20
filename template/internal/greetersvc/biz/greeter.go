package biz

import (
	"<%= go_module_name %>/internal/greetersvc/core/domain"
	"<%= go_module_name %>/internal/greetersvc/core/port"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo port.GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo port.GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *domain.Greeter) (*domain.Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
