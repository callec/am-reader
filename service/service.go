package service

import (
	"context"
	"mag"
	"mag/service/basicservice"
	"mag/service/db"
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
	GetMagazine(context.Context, uuid.UUID) (*mag.Magazine, error)

	// Request a magazine by it's number.
	GetMagazineByNumber(context.Context, int) (*mag.Magazine, error)

	// Get n magazines with m offset.
	ListMagazines(context.Context, int, int) ([]*mag.Magazine, error)

	// Delete some magazine from the database.
	RemoveMagazine(context.Context, uuid.UUID) error

	GetUserByName(context.Context, string) (*mag.User, error)

	GetUser(context.Context, uuid.UUID) (*mag.User, error)

	RegisterUser(context.Context, string, string) error
}

// Factory for db services.
func NewService(r *db.Queries) Service {
	return basicservice.CreateBasicService(r)
}
