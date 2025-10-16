package lessors

import (
	"time"

	"github.com/rrdzgarza/origoDomains/domains/shared"
)

// Lessor represents the primary client offering the SaaS.
type Lessor struct {
	ID        shared.UUID
	Name      string
	Domain    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
