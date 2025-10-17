package authentication

import (
	"time"

	"github.com/rrdzgarza/origo-backend-architecture/domains/shared"
)

// User represents a user account.
type User struct {
	ID        shared.UUID
	LessorID  shared.UUID
	TenantID  shared.UUID
	Email     string
	Phone     string
	Password  string // This would be a hash
	CreatedAt time.Time
	UpdatedAt time.Time
}
