package ports

import (
	"context"

	"github.com/rrdzgarza/origoDomains/domains/lessors"
	"github.com/rrdzgarza/origoDomains/domains/shared"
)

// LessorService defines the port for managing Lessors and their offerings.
type LessorService interface {
	// GetLessorByDomain retrieves a Lessor by its HTTP domain.
	GetLessorByDomain(ctx context.Context, domain string) (*lessors.Lessor, error)

	// GetAppsForLessor retrieves all applications offered by a specific Lessor.
	GetAppsForLessor(ctx context.Context, lessorID shared.UUID) ([]*lessors.App, error)

	// GetPlansForLessor retrieves all plans associated with a specific Lessor.
	GetPlansForLessor(ctx context.Context, lessorID shared.UUID) ([]*lessors.Plan, error)

	// GetAppsForPlan retrieves all applications included in a specific plan.
	GetAppsForPlan(ctx context.Context, planID shared.UUID) ([]*lessors.App, error)

	// GetScopeForAppInPlan retrieves the specific usage scope for an application within a plan.
	GetScopeForAppInPlan(ctx context.Context, planID, appID shared.UUID) (*lessors.Scope, error)
}