package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/zcubbs/mrelay/cmd/server/config"
	"github.com/zcubbs/mrelay/cmd/server/middleware"
	"github.com/zcubbs/mrelay/cmd/server/repository"
	"net/http"
)

type MailHandler struct {
	MailRepository repository.MailRepository
	SmtpConfig     config.SmtpConfig
	AwsSesConfig   config.AwsSesConfig
}

func NewMailHandler(
	mailRepository repository.MailRepository,
	smtpConfig config.SmtpConfig,
	awsSesConfig config.AwsSesConfig,
) *MailHandler {
	return &MailHandler{
		MailRepository: mailRepository,
		SmtpConfig:     smtpConfig,
		AwsSesConfig:   awsSesConfig,
	}
}

func (h *MailHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.With(middleware.Auth).Post("/smtp", h.handleSendSmtpMail)
	return r
}

func (h *MailHandler) handleSendSmtpMail(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "OK")
}
