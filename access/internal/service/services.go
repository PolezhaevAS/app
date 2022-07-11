package service

import (
	"app/access/internal/models"
	"context"
)

func (s *Access) ListServices(ctx context.Context) (
	services []models.Service, err error) {
	services, err = s.db.Services().List(ctx)
	if err != nil {
		return
	}

	return
}
