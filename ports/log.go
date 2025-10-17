package ports

import (
	"context"

	"github.com/rrdzgarza/origo-backend-architecture/domains/logs"
)

// LoggerService defines the port for logging application events.
type LoggerService interface {
	// Log handles the recording of a log entry.
	Log(ctx context.Context, entry logs.Log) error
}