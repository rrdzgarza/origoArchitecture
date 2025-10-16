package ports

import "context"

// FeatureFlagService defines the port for checking the status of feature flags.
// This allows for dynamic feature toggling without redeploying the application.
type FeatureFlagService interface {
	// IsEnabled checks if a feature flag is enabled for a given context.
	// The context can be used to pass user or tenant information for targeted rollouts.
	IsEnabled(ctx context.Context, flagKey string) (bool, error)
}