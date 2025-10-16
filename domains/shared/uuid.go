package shared

import "github.com/google/uuid"

// UUID represents a unique identifier.
type UUID uuid.UUID

// NewUUID generates a new random UUID.
func NewUUID() UUID {
	return UUID(uuid.New())
}

// UUIDFromString parses a string into a UUID.
func UUIDFromString(s string) (UUID, error) {
	id, err := uuid.Parse(s)
	return UUID(id), err
}

// String returns the string representation of the UUID.
func (u UUID) String() string {
	return uuid.UUID(u).String()
}
