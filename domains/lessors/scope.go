package lessors

import "github.com/rrdzgarza/origo-backend-architecture/domains/shared"

// Scope defines usage parameters or limitations for an app within a plan.
type Scope struct {
	AppID       shared.UUID
	Feature     string
	Limit       int
	Description string
}
