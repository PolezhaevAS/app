package service

import (
	"app/auth/internal/models"
	"context"
)

func (s *Auth) List(ctx context.Context) ([]*models.User, error) {
	list, err := s.db.Users().List(ctx)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *Auth) User(ctx context.Context, id uint64) (*models.User, error) {
	user, err := s.db.Users().User(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Auth) Create(ctx context.Context, name, login, password string) error {
	passSHA1 := models.PasswordSHA1(password, s.getAdminSalt())
	err := s.db.Users().Create(ctx, name, login, passSHA1)
	if err != nil {
		return err
	}

	return nil
}

func (s *Auth) Update(ctx context.Context, u *models.User) error {
	u.Password = models.PasswordSHA1(u.Password, s.getAdminSalt())

	err := s.db.Users().Update(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (s *Auth) Delete(ctx context.Context, id uint64) error {
	err := s.db.Users().Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
