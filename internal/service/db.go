// Package service is responsible for database interaction.
package service

import (
	"database/sql"
	"io/ioutil"
	"mag/internal/service/db"
)

func initQueries(d db.DBTX) *db.Queries {
	return db.New(d)
}

// Initialise the database. Note that a database is wrapped using
// sqlc, and must be accessed using a Service.
func initSQL(d *sql.DB) error {
	instructions, err := ioutil.ReadFile("/init.sql")
	if err != nil {
		return err
	}

	create := string(instructions)
	_, err = d.Exec(create)

	return err
}

func setupSQL(d *sql.DB) error {
	instructions, err := ioutil.ReadFile("/setup.sql")
	if err != nil {
		return err
	}

	create := string(instructions)
	_, err = d.Exec(create)

	return err
}
