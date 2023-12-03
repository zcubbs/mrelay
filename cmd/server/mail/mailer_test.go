package mail

import (
	"github.com/stretchr/testify/require"
	"github.com/zcubbs/mrelay/cmd/server/config"
	"testing"
)

func TestSendEmail(t *testing.T) {
	cfg := config.SmtpConfig{
		Username:    "",
		Password:    "",
		FromName:    "test",
		FromAddress: "test@test.email",
		Host:        "localhost",
		Port:        1025,
	}

	mailer := NewDefaultMailer(cfg)

	content := `
	<h1>Test Email</h1>
	<p>This is a test email</p>
	`

	mail := Mail{
		To:            []string{cfg.FromAddress},
		Cc:            []string{},
		Bcc:           []string{},
		Subject:       "Test email",
		Content:       content,
		AttachedFiles: []string{"./testdata/attachement.md"},
	}

	err := mailer.SendMail(mail)
	require.NoError(t, err)
}
