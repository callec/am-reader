package magazine

import (
	"database/sql"
	"io/ioutil"
	"mag/magazine/db"
	"time"

	"github.com/google/uuid"
)

type Magazine struct {
	Id       uuid.UUID
	Date     time.Time
	Number   int
	Location string
}

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
