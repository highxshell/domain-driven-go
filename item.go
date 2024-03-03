package tavern

import "github.com/google/uuid"

// Item is an entity that represents a item in all Domains
type Item struct {
	// ID an identifier of the entity
	ID          uuid.UUID
	Name        string
	Description string
}
