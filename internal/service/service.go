package service

import (
	"database/sql"
	"mag"
	"mag/internal/service/basicservice"
	"mag/internal/service/db"
)

// Factory for db services.
func NewService(r *db.Queries) mag.Service {
	return basicservice.CreateBasicService(r)
}

func InitDB(loc string) (mag.Service, error) {
	d, err := sql.Open("sqlite", loc)
	if err != nil {
		return nil, err
	}
	err = initSQL(d)
	if err != nil {
		return nil, err
	}
	err = setupSQL(d)
	if err != nil {
		return nil, err
	}

	queries := initQueries(d)
	return NewService(queries), nil
}
