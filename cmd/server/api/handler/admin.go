package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

type AdminHandler struct {
}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}

func (h *AdminHandler) Routes() chi.Router {
	r := chi.NewRouter()
	return r
}

func (h *AdminHandler) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "OK")
}
