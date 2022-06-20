package db

import (
	"app/access/internal/database/queries"
	"context"
	"database/sql"
)

var _ Users = (*UsersRepo)(nil)

type UsersRepo struct {
	*sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db}
}

func (r *UsersRepo) Access(ctx context.Context, id uint64) (map[string][]string, error) {

	stmt, err := r.PrepareContext(ctx, queries.USERS_ACCESS)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	access := make(map[string][]string)

	for rows.Next() {
		var service string
		var method string

		err := rows.Scan(&service, &method)
		if err != nil {
			return nil, err
		}

		access[service] = append(access[service], method)
	}

	return access, nil
}

func (r *UsersRepo) Users(ctx context.Context, id uint64) ([]uint64, error) {
	stmt, err := r.PrepareContext(ctx, queries.USERS_LIST)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersId []uint64

	for rows.Next() {
		var userId uint64
		err := rows.Scan(&userId)
		if err != nil {
			return nil, err
		}

		usersId = append(usersId, userId)
	}

	return usersId, nil
}

func (r *UsersRepo) Add(ctx context.Context, id uint64, userId uint64) error {
	stmt, err := r.PrepareContext(ctx, queries.USERS_ADD)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userId, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) Remove(ctx context.Context, id uint64, userId uint64) error {
	stmt, err := r.PrepareContext(ctx, queries.USERS_REMOVE)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id, userId)
	if err != nil {
		return err
	}

	return nil
}
