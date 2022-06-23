package db

import (
	"app/access/internal/database/queries"
	"app/access/internal/models"
	"context"
	"database/sql"
)

var _ Groups = (*GroupsRepo)(nil)

type GroupsRepo struct {
	*sql.DB
}

func NewGroupsRepo(db *sql.DB) *GroupsRepo {
	return &GroupsRepo{db}
}

func (r *GroupsRepo) List(ctx context.Context) ([]*models.Group, error) {
	stmt, err := r.PrepareContext(ctx, queries.GROUP_LIST)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []*models.Group
	for rows.Next() {
		var group models.Group

		err := rows.Scan(&group.ID, &group.Name, &group.Desc)
		if err != nil {
			return nil, err
		}

		group.Users, err = r.listUsers(ctx, group.ID)
		if err != nil {
			return nil, err
		}

		group.Services, err = r.listServices(ctx, group.ID)
		if err != nil {
			return nil, err
		}

		groups = append(groups, &group)
	}

	return groups, nil
}

func (r *GroupsRepo) Group(ctx context.Context, id uint64) (*models.Group, error) {
	stmt, err := r.PrepareContext(ctx, queries.GROUP)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var group *models.Group
	err = stmt.QueryRowContext(ctx, id).Scan(&group.ID, &group.Name, &group.Desc)
	if err != nil {
		return nil, err
	}

	group.Users, err = r.listUsers(ctx, id)
	if err != nil {
		return nil, err
	}

	group.Services, err = r.listServices(ctx, id)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (r *GroupsRepo) listUsers(ctx context.Context, groupId uint64) ([]uint64, error) {
	stmt, err := r.PrepareContext(ctx, queries.USERS_LIST)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, groupId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []uint64
	for rows.Next() {
		var user uint64
		err := rows.Scan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *GroupsRepo) listServices(ctx context.Context, groupId uint64) ([]*models.Service, error) {
	stmt, err := r.PrepareContext(ctx, queries.GROUP_SERVICES)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, groupId)
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

func (r *GroupsRepo) methodsList(ctx context.Context, serviceId uint64) ([]*models.Method, error) {
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

func (r *GroupsRepo) Create(ctx context.Context, name string, desc string) (*models.Group, error) {
	stmt, err := r.PrepareContext(ctx, queries.GROUP_CREATE)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRowContext(ctx, name, desc).Scan(&id)
	if err != nil {
		return nil, err
	}

	group := &models.Group{
		ID:   uint64(id),
		Name: name,
		Desc: desc,
	}

	return group, nil
}

func (r *GroupsRepo) Update(ctx context.Context, group *models.Group) error {
	stmt, err := r.PrepareContext(ctx, queries.GROUP_UPDATE)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, group.ID, group.Name, group.Desc)
	if err != nil {
		return err
	}

	return nil
}

func (r *GroupsRepo) Delete(ctx context.Context, id uint64) error {
	stmt, err := r.PrepareContext(ctx, queries.GROUP_DELETE)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *GroupsRepo) AddMethod(ctx context.Context, id uint64, methodId uint64) error {
	stmt, err := r.PrepareContext(ctx, queries.GROUP_ADD_METHOD)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id, methodId)
	if err != nil {
		return err
	}

	return nil
}

func (r *GroupsRepo) RemoveMethod(ctx context.Context, id uint64, methodId uint64) error {
	stmt, err := r.PrepareContext(ctx, queries.GROUP_REMOVE_METHOD)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id, methodId)
	if err != nil {
		return err
	}

	return nil
}
