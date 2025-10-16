package ports

import (
	"context"

	"github.com/rrdzgarza/origoDomains/domains/smtp"
)

// SMTPService defines the port for sending emails.
type SMTPService interface {
	// SendEmail sends an email message.
	SendEmail(ctx context.Context, email smtp.Email) error
}