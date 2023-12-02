package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

type OpsHandler struct {
	Version string
	Commit  string
	Date    string
}

func NewOpsHandler(version, commit, date string) *OpsHandler {
	return &OpsHandler{
		Version: version,
		Commit:  commit,
		Date:    date,
	}
}

func (h *OpsHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/health", h.Health)
	r.Get("/version", h.VersionInfo)
	return r
}

func (h *OpsHandler) Health(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{"status": "ok"})
}

func (h *OpsHandler) VersionInfo(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{
		"version": h.Version,
		"commit":  h.Commit,
		"date":    h.Date,
	})
}
