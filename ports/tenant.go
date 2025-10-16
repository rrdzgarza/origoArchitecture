package ports

import (
	"context"

	"github.com/rrdzgarza/origoDomains/domains/shared"
	"github.com/rrdzgarza/origoDomains/domains/tenants"
)

// TenantService defines the port for managing tenants and their database instances.
type TenantService interface {
	// GetTenantByID retrieves a tenant by its unique identifier.
	GetTenantByID(ctx context.Context, tenantID shared.UUID) (*tenants.Tenant, error)

	// GetTenantsForLessor retrieves all tenants associated with a specific Lessor.
	GetTenantsForLessor(ctx context.Context, lessorID shared.UUID) ([]*tenants.Tenant, error)

	// GetDBInstancesForTenant retrieves all database instances for a a specific tenant.
	GetDBInstancesForTenant(ctx context.Context, tenantID shared.UUID) ([]*tenants.DatabaseConnection, error)

	// GetDBInstanceForApp retrieves the specific database instance for an application
	// used by a tenant. This assumes a single database can host multiple apps.
	GetDBInstanceForApp(ctx context.Context, tenantID, appID shared.UUID) (*tenants.DatabaseConnection, error)
}