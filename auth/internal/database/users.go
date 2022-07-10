package db

import (
	"context"
	"errors"
	"fmt"

	database "app/internal/sql"

	"app/auth/internal/models"
)

const (
	// Sign in query
	SIGN_IN = `
		SELECT id, name, login, email, password from users.users
		WHERE login = $1 and password = $2
	`

	// list query
	LIST = `
		SELECT id, name, login, email, password from users.users
		WHERE id > $1
		LIMIT $2
	`

	// By id query
	BY_ID = `
		SELECT id, name, login, email, password from users.users
		WHERE id = $1
	`

	// Create query
	CREATE = `
		INSERT INTO users.users(name, login, password)
		VALUES($1, $1, $2)
	`

	// Update query
	UPDATE = `
		UPDATE users.users
		SET
			name = $2,
			login = $3,
			email = $4
		WHERE id = $1
	`

	// Delete query
	DELETE = `
		DELETE FROM users.users WHERE id = $1
	`

	// Get password by id
	// Change password query
	CHANGE_PASSWORD = `
		UPDATE users.users
		SET password = $2
		WHERE id = $1
	`
)

func (db *DB) SignIn(ctx context.Context,
	login, password string) (user models.User, err error) {

	_, err = db.ExecQuery(ctx, database.Get, SIGN_IN, &user,
		login, password)
	if err != nil {
		return
	}

	return
}

func (db *DB) List(ctx context.Context,
	lastID, limit uint64) (users []models.User, err error) {
	if limit > db.MaxLimit() {
		return users, fmt.Errorf("max limit is %d", db.MaxLimit())
	}
	_, err = db.ExecQuery(ctx, database.Select, LIST, &users,
		lastID, limit)
	if err != nil {
		return
	}

	return
}

func (db *DB) ByID(ctx context.Context,
	id uint64) (user models.User, err error) {
	_, err = db.ExecQuery(ctx, database.Get, BY_ID, &user,
		id)
	if err != nil {
		return
	}

	return
}

func (db *DB) Create(ctx context.Context,
	login, password string) (id uint64, err error) {
	id, err = db.ExecQuery(ctx, database.ExecWithReturningId,
		CREATE, nil, login, password)
	if err != nil {
		return
	}
	return
}

func (db *DB) Update(ctx context.Context,
	user models.User) (err error) {
	_, err = db.ExecQuery(ctx, database.Exec, UPDATE, nil,
		user.ID, user.Name, user.Login, user.Email)
	if err != nil {
		return
	}

	return
}

func (db *DB) Delete(ctx context.Context,
	id uint64) (err error) {
	_, err = db.ExecQuery(ctx, database.ExecWithReturningId,
		DELETE, nil, id)
	if err != nil {
		return
	}

	return
}

func (db *DB) ChangePassword(ctx context.Context,
	id uint64, oldPass, newPass string, isReset bool) error {
	var err error
	if !isReset {
		user, err := db.ByID(ctx, id)
		if err != nil {
			return err
		}

		if user.Password != oldPass {
			return errors.New("wrong old password")
		}
	}

	_, err = db.ExecQuery(ctx, database.Exec, CHANGE_PASSWORD, nil,
		id, newPass)
	if err != nil {
		return err
	}

	return nil
}
