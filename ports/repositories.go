package ports

import (
	"context"

	"github.com/rrdzgarza/origoDomains/domains/authentication"
	"github.com/rrdzgarza/origoDomains/domains/lessors"
	"github.com/rrdzgarza/origoDomains/domains/logs"
	"github.com/rrdzgarza/origoDomains/domains/shared"
	"github.com/rrdzgarza/origoDomains/domains/smtp"
	"github.com/rrdzgarza/origoDomains/domains/tenants"
)

// UserRepository defines the persistence port for the User domain.
type UserRepository interface {
	CreateUser(ctx context.Context, user *authentication.User) error
	GetUserByID(ctx context.Context, id shared.UUID) (*authentication.User, error)
	GetUserByEmail(ctx context.Context, email string) (*authentication.User, error)
	UpdateUser(ctx context.Context, user *authentication.User) error
	DeleteUser(ctx context.Context, id shared.UUID) error
}

// LessorRepository defines the persistence port for the Lessor domain.
type LessorRepository interface {
	CreateLessor(ctx context.Context, lessor *lessors.Lessor) error
	GetLessorByID(ctx context.Context, id shared.UUID) (*lessors.Lessor, error)
	GetLessorByDomain(ctx context.Context, domain string) (*lessors.Lessor, error)
	UpdateLessor(ctx context.Context, lessor *lessors.Lessor) error
	DeleteLessor(ctx context.Context, id shared.UUID) error
}

// AppRepository defines the persistence port for the App domain.
type AppRepository interface {
	CreateApp(ctx context.Context, app *lessors.App) error
	GetAppByID(ctx context.Context, id shared.UUID) (*lessors.App, error)
	UpdateApp(ctx context.Context, app *lessors.App) error
	DeleteApp(ctx context.Context, id shared.UUID) error
}

// PlanRepository defines the persistence port for the Plan domain.
type PlanRepository interface {
	CreatePlan(ctx context.Context, plan *lessors.Plan) error
	GetPlanByID(ctx context.Context, id shared.UUID) (*lessors.Plan, error)
	UpdatePlan(ctx context.Context, plan *lessors.Plan) error
	DeletePlan(ctx context.Context, id shared.UUID) error
}

// ScopeRepository defines the persistence port for the Scope domain.
type ScopeRepository interface {
	CreateScope(ctx context.Context, scope *lessors.Scope) error
	GetScopeByID(ctx context.Context, id shared.UUID) (*lessors.Scope, error)
	UpdateScope(ctx context.Context, scope *lessors.Scope) error
	DeleteScope(ctx context.Context, id shared.UUID) error
}

// LogRepository defines the persistence port for the Log domain.
type LogRepository interface {
	CreateLog(ctx context.Context, log *logs.Log) error
}

// EmailRepository defines the persistence port for the Email domain.
type EmailRepository interface {
	CreateEmail(ctx context.Context, email *smtp.Email) error
}

// TenantRepository defines the persistence port for the Tenant domain.
type TenantRepository interface {
	CreateTenant(ctx context.Context, tenant *tenants.Tenant) error
	GetTenantByID(ctx context.Context, id shared.UUID) (*tenants.Tenant, error)
	UpdateTenant(ctx context.Context, tenant *tenants.Tenant) error
	DeleteTenant(ctx context.Context, id shared.UUID) error
}

// DBInstanceRepository defines the persistence port for the DBInstance domain.
type DBInstanceRepository interface {
	CreateDBInstance(ctx context.Context, dbInstance *tenants.DatabaseConnection) error
	GetDBInstanceByID(ctx context.Context, id shared.UUID) (*tenants.DatabaseConnection, error)
	UpdateDBInstance(ctx context.Context, dbInstance *tenants.DatabaseConnection) error
	DeleteDBInstance(ctx context.Context, id shared.UUID) error
}

// InstanceRepository defines the persistence port for the Instance domain.
type InstanceRepository interface {
	CreateInstance(ctx context.Context, instance *tenants.Instance) error
	GetInstanceByID(ctx context.Context, id shared.UUID) (*tenants.Instance, error)
	UpdateInstance(ctx context.Context, instance *tenants.Instance) error
	DeleteInstance(ctx context.Context, id shared.UUID) error
}