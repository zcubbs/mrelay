package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/zcubbs/log"
	"github.com/zcubbs/mrelay/cmd/server/auth"
	"github.com/zcubbs/mrelay/cmd/server/db"
	"github.com/zcubbs/mrelay/cmd/server/mail"
	"github.com/zcubbs/mrelay/cmd/server/models"
	"net/http"
	"strings"
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

type SendMailRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Cc      []string `json:"cc"`
	Bcc     []string `json:"bcc"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
	Account Account  `json:"account"`
}

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// handleSendSmtpMail - Send mail via SMTP.
// @Summary Send mail via SMTP.
// @Description Send mail via SMTP.
// @Tags mail
// @Accept  json
// @Produce  json
// @Param mail body Mail true "mail"
// @Success 200 {string} response "api response"
// @Router /smtp [post]
func (h *MailHandler) handleSendSmtpMail(w http.ResponseWriter, r *http.Request) {
	var req SendMailRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the request
	if req.From == "" || len(req.To) == 0 || req.Subject == "" || req.Body == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Validate the account
	if req.Account.Username == "" || req.Account.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Prepare the mail
	m := mail.Mail{
		To:      req.To,
		Cc:      req.Cc,
		Bcc:     req.Bcc,
		Subject: req.Subject,
		Content: req.Body,
	}

	// Send the mail
	err = h.Mailer.SendMail(m)
	if err != nil {
		http.Error(w, "Failed to send mail", http.StatusInternalServerError)
		return
	}

	// Save the mail
	err = h.MailStore.SaveMail(&models.Email{
		Account:   req.Account.Username,
		ToAddress: strings.Join(req.To, ";"),
		//FromAddress: h.Mailer.GetFromAddress(),
		Subject: m.Subject,
		Body:    m.Content,
		Errors:  "",
	})
	if err != nil {
		log.Error("Failed to save mail",
			"error", err,
			"mail", m,
		)
	}

	// Return a success response
	render.JSON(w, r, "Mail sent successfully")
}
