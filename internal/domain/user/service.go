package user

import (
	"context"
	"errors"
	"time"
	"uuid"

	"github.com/lucaastorres7/karakuri/pkg"
)

type Service interface {
	Register(ctx context.Context, name, email, password, gender string) (*User, error)
	Get(ctx context.Context, id uuid.UUID) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Register(ctx context.Context, name, email, password, gender string) (*User, error) {
	existingUser, _ := s.repo.FindByEmail(ctx, email)
	if existingUser != nil {
		return nil, errors.New("This email is already in use")
	}
	
	hashedPassword, err := pkg.HashPassword(password)
	if err != nil {
		return nil, err
	}

	newUser := &User{
		Id: uuid.NewV7(),
		Name: name,
		Email: email,
		Password: hashedPassword,
		Gender: gender,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, newUser); err != nil {
		return nil, errors.New("Failed to save user in the system")
	}

	return newUser, nil
}

func (s *service) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	u, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, errors.New("User not found")
	}

	return u, nil
}