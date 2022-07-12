package db

import (
	"context"

	sqldb "app/internal/sql"

	"app/access/internal/models"
)

type Repository interface {
	Groups() Groups
	Users() Users
	Services() Services

	FirstStart(context.Context) error
	Close()
}

// Groups repo -
// work with group repository
type Groups interface {
	// List -
	// get page list groups from db by last id gruop and limit
	// Return: list group or error
	List(ctx context.Context,
		lastID, limit uint64) (groups []models.Group, err error)

	// Group -
	// get group by id
	// Return: group or error
	Group(ctx context.Context,
		groupID uint64) (group models.Group, err error)

	// Create -
	// Create group
	// Return: last index or error
	Create(ctx context.Context,
		name, description string) (indexID uint64, err error)

	// Update -
	// update group by id
	// Return: nil or error
	Update(ctx context.Context,
		groupID uint64, name, description string) error

	// Delete -
	// Delete group by id
	// Return: nil or error
	Delete(ctx context.Context,
		groupID uint64) (err error)

	// Add method -
	// add method to group by group id and method id
	// Return: nil or error
	AddMethod(ctx context.Context,
		groupID, methodID uint64) error

	// Remove method -
	// remove method from group by group id and method
	// Return: nil or error
	RemoveMethod(ctx context.Context,
		groupID, methodID uint64) error
}

type Users interface {
	// Access -
	// get user access by user id
	// Return: map[service_name][]method_name or error
	Access(ctx context.Context,
		userID uint64) (map[string][]string, error)

	// Users -
	// get list users id in group by group id
	// Return: list of users id or error
	Users(ctx context.Context,
		groupID uint64) ([]uint64, error)

	// Add -
	// add user into group by group id and user id
	// Return: nil or error
	Add(ctx context.Context,
		groupID uint64, userID uint64) error

	// Remove -
	// remove user from group by group id and user id
	// Return: nil or error
	Remove(ctx context.Context,
		groupID uint64, userID uint64) error
}

type Services interface {
	// List -
	// get list service with methods
	// Return: list of services or error
	List(ctx context.Context) (
		services []models.Service, err error)
}

type DB struct {
	*sqldb.Database
}

func New(cfg *sqldb.Config) (Repository, error) {

	sqlDB, err := sqldb.New(cfg)
	if err != nil {
		return nil, err
	}

	db := new(DB)
	db.Database = sqlDB

	return db, nil
}

func (db *DB) Groups() Groups {
	return NewGroupsRepo(db.Database)
}

func (db *DB) Users() Users {
	return NewUsersRepo(db.Database)
}

func (db *DB) Services() Services {
	return NewServicesRepo(db.Database)
}
