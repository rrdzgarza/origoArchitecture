package ports

import (
	"context"

	"github.com/rrdzgarza/origo-backend-architecture/domains/authentication"
	"github.com/rrdzgarza/origo-backend-architecture/domains/shared"
)

// AuthRequest represents the data for an authentication request.
type AuthRequest struct {
	Strategy  string // e.g., "email", "phone", "google", "facebook"
	Principal string // email, phone number, or provider token
	Secret    string // password, otp, or provider secret
}

// AuthResponse represents the data returned after a successful authentication.
type AuthResponse struct {
	AccessToken  string
	RefreshToken string
	User         authentication.User
}

// MagicLinkRequest represents the data for a magic link request.
type MagicLinkRequest struct {
	Email string
}

// TOTPSetup represents the data for setting up TOTP.
type TOTPSetup struct {
	QRCodeImage string
	Secret      string
}

// AuthenticationService defines the port for authentication-related operations.
type AuthenticationService interface {
	// Login handles authentication with different strategies (email/pass, phone/pass).
	Login(ctx context.Context, req AuthRequest) (*AuthResponse, error)

	// LoginWithProvider handles OAuth 2.0 authentication (Google, Facebook).
	LoginWithProvider(ctx context.Context, provider, token string) (*AuthResponse, error)

	// RequestMagicLink sends a magic link to the user's email for passwordless login.
	RequestMagicLink(ctx context.Context, req MagicLinkRequest) error

	// VerifyMagicLink verifies the token from a magic link.
	VerifyMagicLink(ctx context.Context, token string) (*AuthResponse, error)

	// RefreshToken generates a new access token using a refresh token.
	RefreshToken(ctx context.Context, refreshToken string) (*AuthResponse, error)

	// RevokeToken invalidates a refresh token.
	RevokeToken(ctx context.Context, refreshToken string) error

	// SetupTOTP generates the necessary information for a user to set up 2FA.
	SetupTOTP(ctx context.Context, userID shared.UUID) (*TOTPSetup, error)

	// VerifyTOTP verifies a Time-based One-Time Password.
	VerifyTOTP(ctx context.Context, userID shared.UUID, otp string) (bool, error)
}

// RateLimiter defines the port for rate limiting login attempts.
type RateLimiter interface {
	// IsAllowed checks if a request from a given identifier is allowed.
	// It should return false if the rate limit has been exceeded.
	IsAllowed(ctx context.Context, identifier string) (bool, error)
}

// InputValidator defines the port for validating input data on endpoints.
type InputValidator interface {
	// Validate checks the validity of the given data structure.
	Validate(ctx context.Context, data any) error
}