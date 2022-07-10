package db

import (
	"app/auth/internal/models"
	sqldb "app/internal/sql"
	"context"

	_ "github.com/lib/pq"
)

type Repository interface {
	Users() Users

	// Close connection
	Close()
}

// Users repo -
// Work with users table
type Users interface {
	// Sign in -
	// check user in db by login and password.
	// Return: user or error
	SignIn(ctx context.Context,
		login, password string) (user *models.User, err error)

	// List -
	// get page list users from db by last id user and limit
	// Return: list users or error
	List(ctx context.Context,
		lastID, limit uint64) (users []*models.User, err error)

	// User -
	// get user from db by id
	// Return: user or error
	User(ctx context.Context,
		userID uint64) (user *models.User, err error)

	// Create -
	// create user in db by login and password.
	// Name user = login.
	// Return: error or nil
	Create(ctx context.Context,
		login, password string) error

	// Update -
	// update user in db by user id.
	// Return: error or nil
	Update(ctx context.Context,
		user *models.User) error

	// Delete -
	// delete user from db by user id.
	// Return: error or nil
	Delete(ctx context.Context,
		userID uint64) error

	// Change password -
	// change user password in db by user id.
	// Return: error or nil
	ChangePassword(ctx context.Context,
		userID uint64,
		oldPassword, newPassword string) error
}

var _ Repository = (*DB)(nil)

type DB struct {
	*sqldb.Database
}

func New(cfg *sqldb.Config) (*DB, error) {
	sqlDB, err := sqldb.New(cfg)
	if err != nil {
		return nil, err
	}

	db := new(DB)
	db.Database = sqlDB

	return db, nil
}

func (db *DB) Users() Users {
	return NewUsersRepo(db.Conn)
}
