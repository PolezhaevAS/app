package db

import (
	"app/access/internal/database/queries"
	"app/access/internal/models"
	"context"
	"database/sql"
)

var _ Services = (*ServicesRepo)(nil)

type ServicesRepo struct {
	*sql.DB
}

func NewServicesRepo(db *sql.DB) *ServicesRepo {
	return &ServicesRepo{db}
}

func (r *ServicesRepo) List(ctx context.Context) ([]*models.Service, error) {

	stmt, err := r.PrepareContext(ctx, queries.SERVICES_LIST)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []*models.Service
	for rows.Next() {
		var service models.Service
		err := rows.Scan(&service.ID, &service.Name)
		if err != nil {
			return nil, err
		}

		service.Methods, err = r.methodsList(ctx, service.ID)
		if err != nil {
			return nil, err
		}

		services = append(services, &service)
	}

	return services, nil
}

func (r *ServicesRepo) methodsList(ctx context.Context, serviceId uint64) ([]*models.Method, error) {
	stmt, err := r.PrepareContext(ctx, queries.METHODS_LIST)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, serviceId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var methods []*models.Method
	for rows.Next() {
		var method models.Method
		err = rows.Scan(&method.ID, &method.Name)
		if err != nil {
			return nil, err
		}

		methods = append(methods, &method)
	}

	return methods, nil
}
