package authentication

import (
	"time"

	"github.com/rrdzgarza/origo-backend-architecture/domains/shared"
)

// RefreshToken represents a long-lived, revocable token.
type RefreshToken struct {
	ID        shared.UUID
	UserID    shared.UUID
	Token     string
	ExpiresAt time.Time
	RevokedAt *time.Time
}
