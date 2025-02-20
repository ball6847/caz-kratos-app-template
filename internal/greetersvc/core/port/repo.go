package port

import (
	"app/internal/greetersvc/core/domain"
	"context"
)

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *domain.Greeter) (*domain.Greeter, error)
	Update(context.Context, *domain.Greeter) (*domain.Greeter, error)
	FindByID(context.Context, int64) (*domain.Greeter, error)
	ListByHello(context.Context, string) ([]*domain.Greeter, error)
	ListAll(context.Context) ([]*domain.Greeter, error)
}
