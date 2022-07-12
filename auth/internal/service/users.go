package service

import (
	"context"

	"app/auth/internal/models"
)

func (s *Auth) List(ctx context.Context,
	lastID, limit uint64) (users []models.User, err error) {

	users, err = s.db.List(ctx, lastID, limit)
	if err != nil {
		return
	}

	return
}

func (s *Auth) User(ctx context.Context,
	id uint64) (user models.User, err error) {
	user, err = s.db.ByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (s *Auth) Create(ctx context.Context,
	login, password string) (err error) {
	pass := s.getPasswordSHA1(password)
	_, err = s.db.Create(ctx, login, pass)
	if err != nil {
		return
	}

	return
}

func (s *Auth) Delete(ctx context.Context, id uint64) (err error) {
	err = s.db.Delete(ctx, id)
	if err != nil {
		return
	}

	return
}
