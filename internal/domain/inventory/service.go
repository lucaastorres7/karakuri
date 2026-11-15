package inventory

import (
	"context"
	"uuid"
)

type Service interface {
	CreateProduct(ctx context.Context, name, description, image string, price float64, initialQty int) (*Product, error)
	GetProduct(ctx context.Context, id uuid.UUID) (*Product, error)
	ListCatalog(ctx context.Context) ([]Product, error)
	GetProductsByIds(ctx context.Context, ids []uuid.UUID) ([]Product, error)
	//ReserveStockForOrder(ctx context.Context, orderID uuid.UUID, items []ItemReserve) error
	//ReleaseStockForOrder(ctx context.Context, orderID uuid.UUID, items []ItemReserve) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateProduct(ctx context.Context, name, description, image string, price float64, initialQty int) (*Product, error) {
	return nil, nil
}

func (s *service) GetProduct(ctx context.Context, id uuid.UUID) (*Product, error) {
	return nil, nil
}

func (s *service) ListCatalog(ctx context.Context) ([]Product, error) {
	return nil, nil
}

func (s *service) GetProductsByIds(ctx context.Context, ids []uuid.UUID) ([]Product, error) {
	return nil, nil
}