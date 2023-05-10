package mag

import (
	"time"

	"github.com/google/uuid"
)

// Internal representation of a magazine.
type Magazine struct {
	Id       uuid.UUID
	Date     time.Time // Time created
	Number   int
	Location string // Location of a magazine on disk.
}
