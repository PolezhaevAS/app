package service

import (
	db "app/access/internal/database"
	"app/access/internal/models"
	"context"
)

type Service interface {
	// Access -
	// request user access by user id
	// Return map[ServiceName][]MethodName or error
	Access(ctx context.Context, userID uint64) (
		map[string][]string, error)

	// List -
	// request list groups by last request id and limit
	// Return group list or error
	List(ctx context.Context,
		lastID, limit uint64) (groups []models.Group, err error)

	// Group -
	// request group by group id
	// Return group or error
	Group(ctx context.Context,
		groupID uint64) (group models.Group, err error)

	// Create group -
	// request to create group
	// Return nil or error
	CreateGroup(ctx context.Context,
		name, desc string) (id uint64, err error)

	// Update group -
	// request update group by group id
	// Return nil or error
	UpdateGroup(ctx context.Context,
		groupID uint64, name, desc string) error

	// Delete group -
	// request delete group by group id
	// Return nil or error
	DeleteGroup(ctx context.Context,
		groupID uint64) error

	// Users -
	// request users in group by group id
	// Return []usersID or error
	Users(ctx context.Context,
		groupID uint64) (usersID []uint64, err error)

	// Add user -
	// request add user id into group by group id
	// Return nil or error
	AddUser(ctx context.Context,
		groupID, userID uint64) error

	// Remove user -
	// request remove user from group by group id
	// Return nil or error
	RemoveUser(ctx context.Context,
		groupID, userID uint64) error

	// Add method -
	// request to add method in group by group id
	// Return nil or error
	AddMethod(ctx context.Context,
		groupID, methodID uint64) error

	// Remove method -
	// request remove method from group by group by id
	// Return nil or error
	RemoveMethod(ctx context.Context,
		groupID, methodID uint64) error

	// List servives -
	// request list services
	// Return []Service or error
	ListServices(ctx context.Context) (services []models.Service,
		err error)
}

type Access struct {
	db db.Repository
}

func New(db db.Repository) Service {
	return &Access{db: db}
}
