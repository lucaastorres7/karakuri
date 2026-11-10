package handler

import (
	"net/http"
	"uuid"

	"github.com/lucaastorres7/karakuri/internal/domain/user"
	"github.com/lucaastorres7/karakuri/pkg"
)

type userHandler struct {
	svc user.Service
}

func NewUserHandler(svc user.Service) *userHandler {
	return &userHandler{
		svc: svc,
	}
}

func (h *userHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /users/{id}", h.handleGetUser)
	mux.HandleFunc("POST /register", h.handleRegister)
}

func (h *userHandler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.PathValue("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		pkg.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	u, err := h.svc.Get(ctx, uuid)
	if err != nil {
		pkg.WriteJSON(w, http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
		return
	}

	pkg.WriteJSON(w, http.StatusOK, u)
}

func (h *userHandler) handleRegister(w http.ResponseWriter, r *http.Request	) {
	ctx := r.Context()
	var payload RegisterUserPayload

	if err := pkg.ParseJSON(r, payload); err != nil {
		pkg.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	_, err := h.svc.Register(ctx, payload.Name, payload.Email, payload.Password, payload.Gender)
	if err != nil {
		pkg.WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	pkg.WriteJSON(w, http.StatusCreated, nil)
}