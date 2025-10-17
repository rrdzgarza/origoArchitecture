package authentication

import "github.com/rrdzgarza/origo-backend-architecture/domains/shared"

// OAuth represents an OAuth 2.0 provider.
type OAuth struct {
	UserID       shared.UUID
	Provider     string // "google", "facebook"
	ProviderID   string
	AccessToken  string
	RefreshToken string
}
