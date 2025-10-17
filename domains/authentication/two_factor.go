package authentication

import "github.com/rrdzgarza/origo-backend-architecture/domains/shared"

// TwoFactor represents the 2FA settings for a user.
type TwoFactor struct {
	UserID    shared.UUID
	Type      string // "totp"
	Secret    string
	IsEnabled bool
}
