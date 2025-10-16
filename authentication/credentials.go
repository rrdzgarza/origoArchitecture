package authentication

// Credentials represents the login credentials for a user.
type Credentials struct {
	Type     string // "email", "phone"
	Value    string
	Password string
}
