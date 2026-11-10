package postgres

import (
	"context"
	"errors"
	"uuid"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lucaastorres7/karakuri/internal/domain/user"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) user.Repository {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, u *user.User) error {
	query := `
		INSERT INTO users (
			id, name, email, password, gender, created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(ctx, query, &u.Id, &u.Name, &u.Email, &u.Password, &u.Gender, &u.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) FindById(ctx context.Context, id uuid.UUID) (*user.User, error) {
	query := `
		SELECT id, name, email, password, gender, created_at
		FROM users
		WHERE id = $1
	`

	var u user.User
	err := r.db.QueryRow(ctx, query, id).Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Gender, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // Not found, no critical error
		}
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	query := `
		SELECT id, name, email, password, gender, created_at
		FROM users
		WHERE email = $1
	`

	var u user.User
	err := r.db.QueryRow(ctx, query, email).Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Gender, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}