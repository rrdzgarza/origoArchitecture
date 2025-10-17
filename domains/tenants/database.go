package tenants

import "github.com/rrdzgarza/origo-backend-architecture/domains/shared"

// DatabaseConnection holds the connection parameters for a database.
type DatabaseConnection struct {
	InstanceID shared.UUID
	Driver     string // e.g., "postgres", "mysql"
	DSN        string // Data Source Name
	AppIDs     []shared.UUID
}
