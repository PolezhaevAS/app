package db

import (
	"context"

	"github.com/lib/pq"

	database "app/internal/sql"
)

const (
	USERS_LIST = `
		SELECT user_id
		FROM "access".user_groups WHERE group_id = $1;
	`

	USERS_ADD = `
		INSERT INTO "access".user_groups
		(user_id, group_id)
		VALUES($1, $2);
	`

	USERS_REMOVE = `
		DELETE FROM "access".user_groups
		WHERE group_id = $1 and user_id = $2;
	`

	USERS_ACCESS = `
		select s."name" , array_agg(m."name") methods  from "access".user_groups ug
		join "access".group_methods gm on ug.group_id = gm.group_id 
		join "access".methods m on gm.method_id = m.id 
		join "access".services s on m.service_id = s.id 
		where user_id  = $1
		group by s.id 
	`
)

var _ Users = (*UsersRepo)(nil)

type UsersRepo struct {
	*database.Database
}

func NewUsersRepo(db *database.Database) *UsersRepo {
	return &UsersRepo{db}
}

type tmpAccess struct {
	Name    string
	Methods pq.StringArray
}

func (r *UsersRepo) Access(ctx context.Context,
	id uint64) (map[string][]string, error) {
	var tmpAccess []tmpAccess

	_, err := r.ExecQuery(ctx, database.Select, USERS_ACCESS,
		&tmpAccess, id)
	if err != nil {
		return nil, err
	}

	m := make(map[string][]string)
	for _, tmp := range tmpAccess {
		m[tmp.Name] = tmp.Methods
	}

	return m, nil
}

func (r *UsersRepo) Users(ctx context.Context,
	groupID uint64) (users []uint64, err error) {
	_, err = r.ExecQuery(ctx, database.Select, USERS_LIST,
		&users, groupID)
	if err != nil {
		return
	}

	return
}

func (r *UsersRepo) Add(ctx context.Context,
	id, userID uint64) error {
	_, err := r.ExecQuery(ctx, database.Exec, USERS_ADD,
		nil, id, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) Remove(ctx context.Context,
	id, userID uint64) error {
	_, err := r.ExecQuery(ctx, database.Exec, USERS_REMOVE,
		nil, id, userID)
	if err != nil {
		return err
	}

	return nil
}
