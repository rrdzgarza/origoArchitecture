package ports

import "context"

// ConfigurationService defines the port for accessing application configuration.
// This decouples the core application from the source of configuration (e.g., env vars, files, config server).
type ConfigurationService interface {
	// GetString returns a configuration value as a string.
	GetString(ctx context.Context, key string) (string, error)
	// GetInt returns a configuration value as an integer.
	GetInt(ctx context.Context, key string) (int, error)
	// GetBool returns a configuration value as a boolean.
	GetBool(ctx context.Context, key string) (bool, error)
}