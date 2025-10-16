package authentication

import (
	"time"

	"github.com/rrdzgarza/origoDomains/shared"
)

// MagicLink represents a single-use login link.
type MagicLink struct {
	ID        shared.UUID
	UserID    shared.UUID
	Token     string // The unique, secure token
	ExpiresAt time.Time
	UsedAt    *time.Time
}
