package handler

import (
	"fmt"
	"net/http"
	"uuid"

	"github.com/go-playground/validator/v10"
	"github.com/lucaastorres7/karakuri/internal/domain/user"
	"github.com/lucaastorres7/karakuri/internal/http/middleware"
	"github.com/lucaastorres7/karakuri/pkg"
)

type userHandler struct {
	svc       user.Service
	validator *validator.Validate
}

func NewUserHandler(svc user.Service, v *validator.Validate) *userHandler {
	return &userHandler{
		svc: svc,
		validator: v,
	}
}

func (h *userHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /users/{id}", h.handleGetUser)
	mux.Handle("GET /users/me", middleware.RequireAuth(http.HandlerFunc(h.handleGetMe)))
	mux.HandleFunc("POST /register", h.handleRegister)
	mux.HandleFunc("POST /login", h.handleLogin)
}

func (h *userHandler) handleGetMe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	claims, ok := pkg.GetUserClaims(ctx)
	if !ok {
		pkg.WriteJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "unauthorized",
		})
		return
	}

	u, err := h.svc.Get(ctx, claims.UserId)
	if err != nil {
		pkg.WriteJSON(w, http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
		return
	}

	pkg.WriteJSON(w, http.StatusOK, u)
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

	if err := pkg.ParseJSON(r, &payload); err != nil {
		pkg.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if err := h.validator.Struct(payload); err != nil {
		errs := err.(validator.ValidationErrors)
		pkg.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": fmt.Errorf("invalid payload %v", errs).Error(),
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

func (h *userHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload LoginUserPayload

	if err := pkg.ParseJSON(r, &payload); err != nil {
		pkg.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if err := h.validator.Struct(payload); err != nil {
		errs := err.(validator.ValidationErrors)
		pkg.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": fmt.Errorf("invalid payload %v", errs).Error(),
		})
		return
	}

	token, err := h.svc.Login(ctx, payload.Email, payload.Password)
	if err != nil {
		pkg.WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	pkg.WriteJSON(w, http.StatusOK, map[string]string{
		"token": token,
	})
}