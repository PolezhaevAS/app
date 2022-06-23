package db

import (
	"app/auth/internal/models"
	sqldb "app/internal/sql"
	"context"
)

type Repository interface {
	Users() Users
}

type Users interface {
	SignIn(ctx context.Context, login, password string) (*models.User, error)
	List(ctx context.Context) ([]*models.User, error)
	User(ctx context.Context, id uint64) (*models.User, error)
	Create(ctx context.Context, name, login, password string) error
	Update(ctx context.Context, u *models.User) error
	Delete(ctx context.Context, id uint64) error
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
