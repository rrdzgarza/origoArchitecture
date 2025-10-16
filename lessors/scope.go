package lessors

import "github.com/rrdzgarza/origoDomains/shared"

// Scope defines usage parameters or limitations for an app within a plan.
type Scope struct {
	AppID       shared.UUID
	Feature     string
	Limit       int
	Description string
}
