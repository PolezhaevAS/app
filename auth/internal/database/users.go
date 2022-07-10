package db

import (
	"app/auth/internal/models"
	"context"

	"github.com/jmoiron/sqlx"
)

var (
	SIGN_IN = "select id, name, login, email from users.users"
)

var _ Users = (*UsersRepo)(nil)

type UsersRepo struct {
	*sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) *UsersRepo {
	return &UsersRepo{db}
}

func (r *UsersRepo) SignIn(ctx context.Context,
	login, password string) (*models.User, error) {

	stmt, err := r.PrepareContext(ctx, queries.SIGN_IN)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRowContext(ctx, login, password).
		Scan(&user.ID, &user.Name, &user.Login, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UsersRepo) List(ctx context.Context,
	lastID, limit uint64) ([]*models.User, error) {
	stmt, err := r.PrepareContext(ctx, queries.LIST)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, lastID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Login,
			&user.Email,
		)

		if err != nil {
			return nil, err
		}
		users = append(users, &user)

	}

	return users, nil
}

func (r *UsersRepo) User(ctx context.Context,
	id uint64) (*models.User, error) {

	stmt, err := r.PrepareContext(ctx, queries.USER)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRowContext(ctx, id).
		Scan(&user.ID, &user.Name, &user.Login, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UsersRepo) Create(ctx context.Context,
	login, password string) error {

	stmt, err := r.PrepareContext(ctx, queries.CREATE)
}

func (r *UsersRepo) Create(ctx context.Context, name, login, password string) error {
	stmt, err := r.PrepareContext(ctx, queries.CREATE)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, name, login, password)
	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) Update(ctx context.Context, u *models.User) error {
	stmt, err := r.PrepareContext(ctx, queries.UPDATE)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, u.ID, u.Name, u.Login, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) Delete(ctx context.Context, id uint64) error {
	stmt, err := r.PrepareContext(ctx, queries.DELETE)
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
