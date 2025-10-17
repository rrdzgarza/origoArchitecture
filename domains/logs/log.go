package logs

import (
	"time"

	"github.com/rrdzgarza/origo-backend-architecture/domains/shared"
)

// Log represents an event log.
type Log struct {
	ID        shared.UUID
	LessorID  shared.UUID
	TenantID  shared.UUID
	UserID    *shared.UUID // User might not always be present
	Level     string       // "info", "warn", "error"
	Message   string
	Context   map[string]interface{}
	Timestamp time.Time
}
