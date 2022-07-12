package db

import (
	"context"

	"github.com/jmoiron/sqlx/types"
	"github.com/lib/pq"

	"app/access/internal/models"
	database "app/internal/sql"
)

const (
	GROUP = `
		select g.id, g.name, g.description, json_agg(json_build_object('id', m.id, 'name', m."name")) methods, 
		(
			select 
			array_agg(ug.user_id)
			from access.user_groups ug 
			where group_id  = g.id
		) users from access.groups g	
		left join access.group_methods gm on g.id = gm.group_id
		left join access.methods m on gm.method_id = m.id 
		where g.id = $1
		group by g.id ;
	`

	GROUP_LIST = `
		select g.id, g.name, g.description, json_agg(json_build_object('id', m.id, 'name', m."name")) methods, 
		(
			select 
			array_agg(ug.user_id)
			from access.user_groups ug 
			where group_id  = g.id
		) users from access.groups g	
		left join access.group_methods gm on g.id = gm.group_id
		left join access.methods m on gm.method_id = m.id 
		where g.id > $1
		group by g.id 
		limit $2;
	`

	GROUP_SERVICES = `
		SELECT s.id, s."name" , m.id, m."name"  
		FROM "access".group_methods gm
		JOIN "access".methods m on gm.method_id = m.id 
		JOIN "access".services s on m.service_id = s.id  
		WHERE gm.group_id = $1;
	`

	GROUP_CREATE = `
		INSERT INTO "access"."groups"
		("name", description)
		VALUES($1, $2)
		RETURNING id;
	`

	GROUP_UPDATE = `
		UPDATE "access"."groups"
		SET "name"=$2, description=$3
		WHERE id=$1;
	`

	GROUP_DELETE = `
		DELETE FROM access.groups WHERE id=$1;
	`

	GROUP_ADD_METHOD = `
		INSERT INTO "access".group_methods
		(group_id, method_id)
		VALUES($1, $2);
	`

	GROUP_REMOVE_METHOD = `
		DELETE FROM access.group_methods 
		WHERE group_id = $1 and method_id = $2;
	`
)

var _ Groups = (*GroupsRepo)(nil)

type GroupsRepo struct {
	*database.Database
}

func NewGroupsRepo(db *database.Database) *GroupsRepo {
	return &GroupsRepo{db}
}

type tmpGroup struct {
	ID          uint64
	Name        string
	Description string
	Methods     types.JSONText
	Users       pq.Int64Array
}

func (tmp *tmpGroup) Group() (g models.Group, err error) {
	var m []models.Method
	err = tmp.Methods.Unmarshal(&m)
	if err != nil {
		return
	}

	var usersID []uint64
	for _, id := range tmp.Users {
		usersID = append(usersID, uint64(id))
	}

	return models.Group{
		ID:      tmp.ID,
		Name:    tmp.Name,
		Desc:    tmp.Description,
		Methods: m,
		Users:   usersID,
	}, nil
}

func (r *GroupsRepo) List(ctx context.Context,
	lastID, limit uint64) (groups []models.Group, err error) {
	var tmpGroups []tmpGroup
	_, err = r.ExecQuery(ctx, database.Select,
		GROUP_LIST, &tmpGroups, lastID, limit)
	if err != nil {
		return
	}

	for _, tmp := range tmpGroups {
		var group models.Group
		group, err = tmp.Group()
		if err != nil {
			return
		}
		groups = append(groups, group)
	}

	return
}

func (r *GroupsRepo) Group(ctx context.Context,
	id uint64) (group models.Group, err error) {
	var tmpGroup tmpGroup
	_, err = r.ExecQuery(ctx, database.Get, GROUP,
		&tmpGroup, id)
	if err != nil {
		return
	}

	return tmpGroup.Group()
}

func (r *GroupsRepo) Create(ctx context.Context,
	name, desc string) (id uint64, err error) {
	id, err = r.ExecQuery(ctx, database.ExecWithReturningId,
		GROUP_CREATE, nil, name, desc)
	if err != nil {
		return
	}

	return
}

func (r *GroupsRepo) Update(ctx context.Context,
	id uint64, name, desc string) error {
	_, err := r.ExecQuery(ctx, database.Exec, GROUP_UPDATE,
		nil, id, name, desc)
	if err != nil {
		return err
	}

	return nil
}

func (r *GroupsRepo) Delete(ctx context.Context, id uint64) error {
	_, err := r.ExecQuery(ctx, database.Exec, GROUP_DELETE,
		nil, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *GroupsRepo) AddMethod(ctx context.Context,
	id, mID uint64) error {
	_, err := r.ExecQuery(ctx, database.Exec, GROUP_ADD_METHOD,
		nil, id, mID)
	if err != nil {
		return err
	}

	return nil
}

func (r *GroupsRepo) RemoveMethod(ctx context.Context,
	id, mID uint64) error {
	_, err := r.ExecQuery(ctx, database.Exec, GROUP_REMOVE_METHOD,
		nil, id, mID)
	if err != nil {
		return err
	}

	return nil
}
