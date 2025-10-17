package lessors

import (
	"time"

	"github.com/rrdzgarza/origo-backend-architecture/domains/shared"
)

// Plan represents a SaaS plan offered by a Lessor.
type Plan struct {
	ID        shared.UUID
	LessorID  shared.UUID
	Name      string
	Price     float64
	AppIDs    []shared.UUID // The applications included in this plan
	CreatedAt time.Time
	UpdatedAt time.Time
}
