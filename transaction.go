package tavern

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is a valueobject because it has no identificator
type Transaction struct {
	// all values lowercase since they are immutable
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
