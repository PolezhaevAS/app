package service

import (
	db "app/access/internal/database"
	"app/access/internal/models"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	// Get user access by user id
	Access(ctx context.Context, userId uint64) (map[string][]string, error)
	// Get list groups
	List(ctx context.Context) ([]*models.Group, error)
	// Get group by id
	Group(ctx context.Context, id uint64) (*models.Group, error)
	// Create new group
	CreateGroup(ctx context.Context, name, descr string) (*models.Group, error)
	// Update group
	UpdateGroup(ctx context.Context, g *models.Group) error
	// Delete group by id
	DeleteGroup(ctx context.Context, id uint64) error
	// Get users in group by group id
	Users(ctx context.Context, id uint64) ([]uint64, error)
	// Add user into group
	AddUser(ctx context.Context, id uint64, userId uint64) error
	// Remove user from group
	RemoveUser(ctx context.Context, id uint64, userId uint64) error
	// Get list services
	ListService(ctx context.Context) ([]*models.Service, error)
	// Add method to group
	AddMethod(ctx context.Context, id uint64, methodId uint64) error
	// Remove method from group
	RemoveMethod(ctx context.Context, id uint64, methodId uint64) error
}

var _ Service = (*Access)(nil)

type Access struct {
	db db.Repository
}

func New(db db.Repository) *Access {
	return &Access{db: db}
}

func (s *Access) Access(ctx context.Context, userId uint64) (map[string][]string, error) {

	if userId <= 0 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid user id: %d", userId))
	}

	access, err := s.db.Users().Access(ctx, userId)
	if err != nil {
		return nil, status.Error(codes.Aborted, fmt.Sprintf("Access error: %s", err.Error()))
	}

	return access, nil
}

func (s *Access) List(ctx context.Context) ([]*models.Group, error) {
	groups, err := s.db.Groups().List(ctx)
	if err != nil {
		return nil, status.Error(codes.Aborted, fmt.Sprintf("List error: %s", err.Error()))
	}

	return groups, nil
}

func (s *Access) Group(ctx context.Context, id uint64) (*models.Group, error) {

	err := s.validateInputGroupId(ctx, id)
	if err != nil {
		return nil, err
	}

	group, err := s.db.Groups().Group(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Aborted, fmt.Sprintf("Group error: %s", err.Error()))
	}

	return group, nil
}

func (s *Access) CreateGroup(ctx context.Context, name, descr string) (*models.Group, error) {

	if name == "" {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid name group: %s. Name group empty", name))
	}

	group, err := s.db.Groups().Create(ctx, name, descr)
	if err != nil {
		return nil, status.Error(codes.Aborted, fmt.Sprintf("Create group error: %s", err.Error()))
	}

	return group, nil
}

func (s *Access) UpdateGroup(ctx context.Context, g *models.Group) error {

	err := s.validateInputGroupId(ctx, g.ID)
	if err != nil {
		return err
	}

	err = s.db.Groups().Update(ctx, g)
	if err != nil {
		return status.Error(codes.Aborted, fmt.Sprintf("Update group error: %s", err.Error()))
	}

	return nil
}

func (s *Access) DeleteGroup(ctx context.Context, id uint64) error {

	err := s.validateInputGroupId(ctx, id)
	if err != nil {
		return err
	}

	err = s.db.Groups().Delete(ctx, id)
	if err != nil {
		return status.Error(codes.Aborted, fmt.Sprintf("Delete group error: %s", err.Error()))
	}

	return nil
}

func (s *Access) Users(ctx context.Context, id uint64) ([]uint64, error) {

	err := s.validateInputGroupId(ctx, id)
	if err != nil {
		return nil, err
	}

	usersId, err := s.db.Users().Users(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Aborted, fmt.Sprintf("Users error: %s", err.Error()))
	}

	return usersId, nil
}

func (s *Access) AddUser(ctx context.Context, id uint64, userId uint64) error {

	err := s.validateInputGroupId(ctx, id)
	if err != nil {
		return err
	}

	err = s.db.Users().Add(ctx, id, userId)
	if err != nil {
		return status.Error(codes.Aborted, fmt.Sprintf("Add user error: %s", err.Error()))
	}

	return nil
}

func (s *Access) RemoveUser(ctx context.Context, id uint64, userId uint64) error {

	err := s.validateInputGroupId(ctx, id)
	if err != nil {
		return err
	}

	err = s.db.Users().Remove(ctx, id, userId)
	if err != nil {
		return status.Error(codes.Aborted, fmt.Sprintf("Remove user error: %s", err.Error()))
	}

	return nil
}

func (s *Access) ListService(ctx context.Context) ([]*models.Service, error) {
	services, err := s.db.Services().List(ctx)
	if err != nil {
		return nil, status.Error(codes.Aborted, fmt.Sprintf("List services error: %s", err.Error()))
	}

	return services, nil
}

func (s *Access) AddMethod(ctx context.Context, id uint64, methodId uint64) error {

	err := s.validateInputGroupId(ctx, id)
	if err != nil {
		return err
	}

	err = s.db.Groups().AddMethod(ctx, id, methodId)
	if err != nil {
		return status.Error(codes.Aborted, fmt.Sprintf("Add method error: %s", err.Error()))
	}

	return nil
}

func (s *Access) RemoveMethod(ctx context.Context, id uint64, methodId uint64) error {

	err := s.validateInputGroupId(ctx, id)
	if err != nil {
		return err
	}

	err = s.db.Groups().RemoveMethod(ctx, id, methodId)
	if err != nil {
		return status.Error(codes.Aborted, fmt.Sprintf("Remove method error: %s", err.Error()))
	}

	return nil
}

func (s *Access) validateInputGroupId(ctx context.Context, id uint64) error {

	if id <= 0 {
		return status.Error(codes.InvalidArgument, fmt.Sprintf("invalid group id: %d", id))
	}

	err := s.checkGroup(ctx, id)
	if err != nil {
		return status.Error(codes.InvalidArgument, fmt.Sprintf("Check exists group: %s", err.Error()))
	}

	return nil
}

func (s *Access) checkGroup(ctx context.Context, id uint64) error {
	g, err := s.db.Groups().Group(ctx, id)
	if err != nil {
		return err
	}

	if g == nil {
		return fmt.Errorf("not exists group with id %d", id)
	}

	return nil
}
