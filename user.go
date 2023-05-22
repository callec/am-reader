package mag

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID
	Username     string
	PasswordHash string // bcrypt hash of the password
	Created      time.Time
	LastOnline   time.Time
}
