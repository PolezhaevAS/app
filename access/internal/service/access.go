package service

import "context"

func (s *Access) Access(ctx context.Context,
	userID uint64) (map[string][]string, error) {
	access, err := s.db.Users().Access(ctx, userID)
	if err != nil {
		return nil, err
	}

	return access, nil
}
