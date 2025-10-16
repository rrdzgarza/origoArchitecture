package lessors

import (
	"time"

	"github.com/rrdzgarza/origoDomains/shared"
)

// Lessor represents the primary client offering the SaaS.
type Lessor struct {
	ID        shared.UUID
	Name      string
	Domain    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
