package postgres

import (
	"context"
	"errors"
	"uuid"

	"github.com/jackc/pgx/v5"
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
	query := `
		INSERT INTO products (
			id, name, description, image, price, quantity, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.db.Exec(ctx, query, &p.Id, &p.Name, &p.Description, &p.Image, &p.Price, &p.Quantity, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *inventoryRepo) GetById(ctx context.Context, id uuid.UUID) (*inventory.Product, error) {
	query := `
		SELECT id, name, description, image, price, quantity, created_at, updated_at
		FROM products
		WHERE id = $1
	`

	var p inventory.Product

	err := r.db.QueryRow(ctx, query, id).Scan(&p.Id, &p.Name, &p.Description, &p.Image, &p.Price, &p.Quantity, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // Not found, no critical error
		}
		return nil, err
	}
	return &p, nil
}

func (r *inventoryRepo) GetByIds(ctx context.Context, ids []uuid.UUID) ([]inventory.Product, error) {
	// If no IDs were provided
	if len(ids) == 0 {
		return []inventory.Product{}, nil
	}

	query := `
		SELECT id, name, description, image, price, quantity, created_at, updated_at
		FROM products
		WHERE id = ANY($1)
	`

	rows, err := r.db.Query(ctx, query, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []inventory.Product

	for rows.Next() {
		var p inventory.Product
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Image, &p.Price, &p.Quantity, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *inventoryRepo) ListAll(ctx context.Context) ([]inventory.Product, error) {
	query := `
		SELECT id, name, description, image, price, quantity, created_at, updated_at
		FROM products
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []inventory.Product

	for rows.Next() {
		var p inventory.Product
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Image, &p.Price, &p.Quantity, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}