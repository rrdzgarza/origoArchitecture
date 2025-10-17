package lessors

import (
	"time"

	"github.com/rrdzgarza/origo-backend-architecture/domains/shared"
)

// App represents an application that a Lessor can offer.
type App struct {
	ID          shared.UUID
	Name        string
	Description string
	CreatedAt   time.Time
}
