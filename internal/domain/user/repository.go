package user

import (
	"context"
	"uuid"
)

type Repository interface {
	Create(ctx context.Context, u *User) error
	FindById(ctx context.Context, id uuid.UUID) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}