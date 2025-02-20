package data

import (
	"context"

	"<%= go_module_name %>/internal/greetersvc/core/domain"
	"<%= go_module_name %>/internal/greetersvc/core/port"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	log *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(logger log.Logger) port.GreeterRepo {
	return &greeterRepo{
		log: log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *domain.Greeter) (*domain.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *domain.Greeter) (*domain.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*domain.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*domain.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*domain.Greeter, error) {
	return nil, nil
}
