package inventory

import (
	"context"
	"uuid"
)

type Repository interface {
	Create(ctx context.Context, p *Product) error
	GetById(ctx context.Context, id uuid.UUID) (*Product, error)
	GetByIds(ctx context.Context, ids []uuid.UUID) ([]Product, error)
	ListAll(ctx context.Context) ([]Product, error)
	//ReserveItems(ctx context.Context, orderID uuid.UUID, items []ItemReserve) error
	//ReleaseItems(ctx context.Context, orderID uuid.UUID, items []ItemReserve) error
}