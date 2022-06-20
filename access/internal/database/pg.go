package db

import (
	"app/access/internal/models"
	sqldb "app/internal/sql"
	"context"

	_ "github.com/lib/pq"
)

type Repository interface {
	FirstStart(context.Context) error
	Close()

	Groups() Groups
	Users() Users
	Services() Services
}

type Groups interface {
	// Get list groups
	List(context.Context) ([]*models.Group, error)

	// Get group by id
	Group(ctx context.Context, id uint64) (*models.Group, error)

	// Create new group
	Create(ctx context.Context, name string, descr string) (*models.Group, error)

	// Update group by id
	Update(ctx context.Context, group *models.Group) error

	// Delete group by id
	Delete(ctx context.Context, id uint64) error

	// Add method into group by group id and method id
	AddMethod(ctx context.Context, id uint64, methodId uint64) error

	// Remove method from group by group id nad method id
	RemoveMethod(ctx context.Context, id uint64, methodId uint64) error
}

type Users interface {
	// Get user access by id
	// Return map[service_name][]method_name
	Access(ctx context.Context, id uint64) (map[string][]string, error)

	// Get users id in group by group id
	Users(ctx context.Context, id uint64) ([]uint64, error)

	// Add user into group by group id and user id
	Add(ctx context.Context, id uint64, userId uint64) error

	// Remove user from group by group id and user id
	Remove(ctx context.Context, id uint64, userId uint64) error
}

type Services interface {
	// Get list service with methods
	List(ctx context.Context) ([]*models.Service, error)
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

func (db *DB) Groups() Groups {
	return NewGroupsRepo(db.Conn)
}

func (db *DB) Services() Services {
	return NewServicesRepo(db.Conn)
}

func (db *DB) Users() Users {
	return NewUsersRepo(db.Conn)
}
