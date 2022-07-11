package db

import (
	"context"
	"database/sql"

	"app/internal/service"
	database "app/internal/sql"
)

const (
	SERVICE_BY_ID = `
		SELECT id FROM "access".services WHERE id = $1 
	`

	SERVICE_ADD = `
		INSERT INTO "access".services
		(id, "name")
		VALUES($1, $2);
	`

	METHOD_BY_NAME = `
		SELECT id FROM "access".methods 
		WHERE service_id = $1 and "name" = $2
	`

	METHOD_ADD = `
		INSERT INTO "access".methods
		(service_id, "name")
		VALUES($1, $2);
	`
)

func (db *DB) FirstStart(ctx context.Context) error {
	err := db.initServices(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) initServices(ctx context.Context) error {
	services := service.Services()

	for pos, service := range services {
		id := pos + 1
		err := db.addService(ctx, id, service.ServiceName)
		if err != nil {
			return err
		}

		for _, method := range service.Methods {
			err = db.addMethods(ctx, id, method.MethodName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (db *DB) addService(ctx context.Context,
	id int, name string) error {
	_, err := db.ExecQuery(ctx, database.Get, SERVICE_BY_ID,
		nil, id)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err := db.ExecQuery(ctx, database.Exec,
				SERVICE_ADD, id, name)
			if err != nil {
				return err
			}

			return nil
		}

		return err
	}

	return nil
}

func (db *DB) addMethods(ctx context.Context,
	serviceID int, name string) error {
	_, err := db.ExecQuery(ctx, database.Get, METHOD_BY_NAME,
		nil, serviceID, name)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err := db.ExecQuery(ctx, database.Exec,
				METHOD_ADD, nil, serviceID, name)
			if err != nil {
				return err
			}

			return nil
		}

		return err
	}

	return nil
}
