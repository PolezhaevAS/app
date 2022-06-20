package db

import (
	"app/access/internal/database/queries"
	"app/internal/service"
	"context"
	"database/sql"
)

func (db *DB) FirstStart(ctx context.Context) error {
	err := db.initServices(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) initServices(ctx context.Context) error {
	services := service.All()

	stmt, err := db.Conn.PrepareContext(ctx, queries.SERVICE_ADD)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for idSlice, service := range services.Services {
		id := idSlice + 1
		// check for exists
		exists, err := db.existsService(ctx, id)
		if err != nil {
			return err
		}

		if !exists {
			// insert service

			_, err = stmt.ExecContext(ctx, id, service.Name)
			if err != nil {
				return err
			}
		}

		// insert methods
		for method := range service.Methods {

			exists, err = db.existsMethod(ctx, id, method)
			if err != nil {
				return err
			}

			if !exists {
				err = db.initMethod(ctx, id, method)
				if err != nil {
					return err
				}
			}
		}

	}

	return nil
}

func (db *DB) existsService(ctx context.Context, idService int) (bool, error) {
	stmt, err := db.Conn.PrepareContext(ctx, queries.SERVICE_EXISTS)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRowContext(ctx, idService).Scan(&id)
	switch err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}

func (db *DB) initMethod(ctx context.Context, idService int, name string) error {
	stmt, err := db.Conn.PrepareContext(ctx, queries.METHOD_ADD)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, idService, name)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) existsMethod(ctx context.Context, idService int, name string) (bool, error) {
	stmt, err := db.Conn.PrepareContext(ctx, queries.METHOD_EXISTS)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRowContext(ctx, idService, name).Scan(&id)
	switch err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}
