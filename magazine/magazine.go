// Package magazine is responsible for representing a magazine
// internally and for database interaction.
//
// In addition to the Magazine struct to represent individual magazines
// you should also use the Service struct for database interaction.
package magazine

import (
	"database/sql"
	"io/ioutil"
	"mag/magazine/db"
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

// Initialise the database. Note that a database is wrapped using
// sqlc, and must be accessed using a Service.
func InitSQL(d *sql.DB) error {
	instructions, err := ioutil.ReadFile("./magazine/init.sql")
	if err != nil {
		return err
	}

	create := string(instructions)
	_, err = d.Exec(create)

	return err
}

// Factory for db services.
func NewService(r *db.Queries) Service {
	return createBasicService(r)
}
