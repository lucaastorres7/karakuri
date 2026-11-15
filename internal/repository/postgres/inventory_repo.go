package postgres

import (
	"context"
	"uuid"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lucaastorres7/karakuri/internal/domain/inventory"
)

type inventoryRepo struct {
	db *pgxpool.Pool
}

func NewInventoryRepository(db *pgxpool.Pool) inventory.Repository{
	return &inventoryRepo{
		db: db,
	}
}

func (r *inventoryRepo) Create(ctx context.Context, p *inventory.Product) error {
	return nil
}

func (r *inventoryRepo) GetById(ctx context.Context, id uuid.UUID) (*inventory.Product, error) {
	return nil, nil
}

func (r *inventoryRepo) GetByIds(ctx context.Context, ids []uuid.UUID) ([]inventory.Product, error) {
	return nil, nil
}

func (r *inventoryRepo) ListAll(ctx context.Context) ([]inventory.Product, error) {
	return nil, nil
}