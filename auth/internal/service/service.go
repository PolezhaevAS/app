package service

import (
	"app/auth/internal/models"
	"context"
)

type Service interface {
	Token(ctx context.Context, login, password string) (string, *models.User, error)
	List(ctx context.Context) ([]*models.User, error)
	User(ctx context.Context, id uint64) (*models.User, error)
	Create(ctx context.Context, name, login, password string) error
	Update(ctx context.Context, u *models.User) error
	Delete(ctx context.Context, id uint64) error
}
