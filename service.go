package mag

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// TODO: factor out to internal package
// Any struct for database interaction must implement these functions.
type Service interface {
	// To add a magazine you must supply the number, date of creation,
	// and the path.
	AddMagazine(context.Context, int, time.Time, string) error

	// Request a magazine by it's id (uuid.UUID).
	GetMagazine(context.Context, uuid.UUID) (*Magazine, error)

	// Request a magazine by it's number.
	GetMagazineByNumber(context.Context, int) (*Magazine, error)

	// Get n magazines with m offset.
	ListMagazines(context.Context, int, int) ([]*Magazine, error)

	// Delete some magazine from the database.
	RemoveMagazine(context.Context, uuid.UUID) error

	GetUserByName(context.Context, string) (*User, error)

	GetUser(context.Context, uuid.UUID) (*User, error)

	RegisterUser(context.Context, string, string) error
}
