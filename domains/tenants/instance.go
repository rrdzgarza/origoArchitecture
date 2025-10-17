package tenants

import "github.com/rrdzgarza/origo-backend-architecture/domains/shared"

// Instance represents a deployment or environment for a Tenant.
type Instance struct {
	ID       shared.UUID
	TenantID shared.UUID
	Name     string
}
