package tenants

import (
	"time"

	"github.com/rrdzgarza/origoDomains/shared"
)

// Tenant represents an end-customer account derived from a Lessor.
type Tenant struct {
	ID        shared.UUID
	LessorID  shared.UUID
	Name      string
	PlanID    shared.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
