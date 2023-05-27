package service

import (
	"mag"
	"mag/internal/service/basicservice"
	"mag/internal/service/db"
)

// Factory for db services.
func NewService(r *db.Queries) mag.Service {
	return basicservice.CreateBasicService(r)
}
