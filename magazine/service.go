package magazine

import (
	"context"
	"mag/magazine/db"
	"time"

	"github.com/google/uuid"
)

// Any service for database interaction must implement these functions.
type Service interface {
	// To add a magazine you must supply the number, date of creation,
	// and the path.
	AddMagazine(context.Context, int, time.Time, string) error

	// Request a magazine by its id (uuid.UUID).
	GetMagazine(context.Context, uuid.UUID) (*Magazine, error)

	// Request a magazine by its number.
	GetMagazineByNumber(context.Context, int) (*Magazine, error)

	// Get n magazines with m offset.
	ListMagazines(context.Context, int, int) ([]*Magazine, error)

	// Delete some magazine from the database.
	removeMagazine(context.Context, uuid.UUID) error
}

func createBasicService(r *db.Queries) *basicService {
	return &basicService{
		r: r,
	}
}
