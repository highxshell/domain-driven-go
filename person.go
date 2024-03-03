// Package entities holds all the entities that are shared across all subdomains
package tavern

import "github.com/google/uuid"

// Person is an entity that represents a person in all Domains
type Person struct {
	// ID an identifier of the entity
	ID   uuid.UUID
	Name string
	Age  int
}
