package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/login", h.handleLogin)
	r.Post("/register", h.handleRegister)
	r.Post("/logout", h.handleLogout)
	r.Post("/refresh", h.handleRefresh)
	return r
}

func (h *AuthHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "unimplemented")
}

func (h *AuthHandler) handleRegister(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "unimplemented")
}

func (h *AuthHandler) handleLogout(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "unimplemented")
}

func (h *AuthHandler) handleRefresh(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "unimplemented")
}
