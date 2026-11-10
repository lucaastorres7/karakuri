package handler

import "github.com/lucaastorres7/karakuri/internal/domain/user"

type userHandler struct {
	svc user.Service
}

func NewUserHandler(svc user.Service) *userHandler {
	return &userHandler{
		svc: svc,
	}
}

func (h *userHandler) 