package smtp

import (
	"time"

	"github.com/rrdzgarza/origoArchitecture/domains/shared"
)

// Email represents an email to be sent.
type Email struct {
	ID        shared.UUID
	LessorID  shared.UUID
	TenantID  shared.UUID
	From      string
	To        []string
	Subject   string
	Body      string
	SentAt    *time.Time
	CreatedAt time.Time
}
