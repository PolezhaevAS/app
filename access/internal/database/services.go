package db

import (
	"context"

	"github.com/jmoiron/sqlx/types"

	"app/access/internal/models"
	database "app/internal/sql"
)

const (
	SERVICE_LIST = `
		select s.id , s."name" , json_agg(json_build_object('id', m.id, 'name', m."name")) methods from access.methods m 
		join access.services s on m.service_id = s.id 
		group by s.id 
	`
)

var _ Services = (*ServicesRepo)(nil)

type ServicesRepo struct {
	*database.Database
}

func NewServicesRepo(db *database.Database) *ServicesRepo {
	return &ServicesRepo{db}
}

type tmpService struct {
	ID      uint64
	Name    string
	Methods types.JSONText
}

func (tmp *tmpService) Service() (s models.Service, err error) {
	var m []models.Method
	err = tmp.Methods.Unmarshal(&m)
	if err != nil {
		return
	}

	return models.Service{
		ID:      tmp.ID,
		Name:    tmp.Name,
		Methods: m,
	}, nil
}

func (r *ServicesRepo) List(ctx context.Context) (
	services []models.Service, err error) {
	var tmpServices []tmpService
	_, err = r.ExecQuery(ctx, database.Select, SERVICE_LIST,
		&tmpServices)
	if err != nil {
		return
	}

	for _, tmp := range tmpServices {
		var service models.Service
		service, err = tmp.Service()
		if err != nil {
			return
		}
		services = append(services, service)
	}

	return
}
