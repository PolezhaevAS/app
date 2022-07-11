package service

import (
	"app/access/internal/models"
	"context"
)

func (s *Access) List(ctx context.Context,
	lastID, limit uint64) (groups []models.Group, err error) {
	groups, err = s.db.Groups().List(ctx, lastID, limit)
	if err != nil {
		return
	}

	return
}

func (s *Access) Group(ctx context.Context,
	groupID uint64) (group models.Group, err error) {
	group, err = s.db.Groups().Group(ctx, groupID)
	if err != nil {
		return
	}

	return
}

func (s *Access) CreateGroup(ctx context.Context,
	name, desc string) (id uint64, err error) {
	id, err = s.db.Groups().Create(ctx, name, desc)
	if err != nil {
		return
	}

	return
}

func (s *Access) UpdateGroup(ctx context.Context,
	groupID uint64, name, desc string) error {
	err := s.db.Groups().Update(ctx, groupID, name, desc)
	if err != nil {
		return err
	}

	return nil
}

func (s *Access) DeleteGroup(ctx context.Context,
	groupID uint64) error {
	err := s.db.Groups().Delete(ctx, groupID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Access) AddMethod(ctx context.Context,
	groupID, methodID uint64) error {
	err := s.db.Groups().AddMethod(ctx, groupID, methodID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Access) RemoveMethod(ctx context.Context,
	groupID, methodID uint64) error {
	err := s.db.Groups().RemoveMethod(ctx, groupID, methodID)
	if err != nil {
		return err
	}

	return nil
}
