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

		err := rows.Scan(&group.ID, &group.Name, &group.Descr)
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
	err = stmt.QueryRowContext(ctx, id).Scan(&group.ID, &group.Name, &group.Descr)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (r *GroupsRepo) Create(ctx context.Context, name string, descr string) (*models.Group, error) {
	stmt, err := r.PrepareContext(ctx, queries.GROUP_CREATE)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRowContext(ctx, name, descr).Scan(&id)
	if err != nil {
		return nil, err
	}

	group := &models.Group{
		ID:    uint64(id),
		Name:  name,
		Descr: descr,
	}

	return group, nil
}

func (r *GroupsRepo) Update(ctx context.Context, group *models.Group) error {
	stmt, err := r.PrepareContext(ctx, queries.GROUP_UPDATE)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, group.ID, group.Name, group.Descr)
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
