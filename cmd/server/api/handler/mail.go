package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/zcubbs/mrelay/cmd/server/auth"
	"github.com/zcubbs/mrelay/cmd/server/db"
	"github.com/zcubbs/mrelay/cmd/server/mail"
	"net/http"
)

type MailHandler struct {
	MailStore db.MailStore
	Mailer    mail.Mailer
}

func NewMailHandler(
	mailStore db.MailStore,
	mailer mail.Mailer,
) *MailHandler {
	return &MailHandler{
		MailStore: mailStore,
		Mailer:    mailer,
	}
}

func (h *MailHandler) Routes(realm string, store db.AuthStore) chi.Router {
	r := chi.NewRouter()
	r.With(auth.BasicAuth(realm, store)).Post("/smtp", h.handleSendSmtpMail)
	return r
}

func (h *MailHandler) handleSendSmtpMail(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "OK")
}
