package service

import "context"

func (s *Access) Users(ctx context.Context,
	groupID uint64) (usersID []uint64, err error) {
	usersID, err = s.db.Users().Users(ctx, groupID)
	if err != nil {
		return
	}

	return
}

func (s *Access) AddUser(ctx context.Context,
	groupID, userID uint64) error {
	err := s.db.Users().Add(ctx, groupID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Access) RemoveUser(ctx context.Context,
	groupID, userID uint64) error {
	err := s.db.Users().Remove(ctx, groupID, userID)
	if err != nil {
		return err
	}

	return nil
}
