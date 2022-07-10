package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"app/internal/token"

	"app/access/pkg/access"

	"app/auth/internal/config"
	db "app/auth/internal/database"
	"app/auth/internal/models"
)

// Service - work with users and authorization
type Service interface {

	// Sign in -
	// request token by login and password.
	// Return token and user model or error
	SignIn(ctx context.Context,
		login, password string) (token string,
		user models.User, access map[string][]string,
		err error)

	// List -
	// request list users by last user id
	// (Using in query WHERE id > $1)
	// and limit.
	// Return user list
	List(ctx context.Context,
		lastID, limit uint64) (users []models.User, err error)

	// User -
	// request user by id.
	// Return user model
	User(ctx context.Context,
		userID uint64) (user models.User, err error)

	// Create -
	// request to create new user with login and password.
	// User name = login.
	// Return error or nil
	Create(ctx context.Context, login, password string) error

	// Delete -
	// request to delete user by id.
	// Return error or nil
	Delete(ctx context.Context, userID uint64) error

	// Change user -
	// request to change user name/login/email by id from claims.
	// Return error or nil
	ChangeUser(ctx context.Context,
		name, login, email string) error

	// Change user password -
	// request to change user password by id from claims.
	// Return error or nil
	ChangeUserPassword(ctx context.Context,
		oldPassword, newPassword string, isReset bool) error
}

type Auth struct {
	db     db.Repository
	cfg    *config.Config
	m      *token.Source
	access *access.Client
}

func New(db db.Repository,
	cfg *config.Config, m *token.Source,
	access *access.Client) Service {
	return &Auth{
		db:     db,
		cfg:    cfg,
		m:      m,
		access: access,
	}
}

func (s *Auth) getAdminSalt() string {
	return s.cfg.AuthConfig.Salt
}

func (s *Auth) getPasswordSHA1(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x",
		hash.Sum([]byte(s.getAdminSalt())))
}

func (s *Auth) getExp() int64 {
	return time.
		Now().
		Add(time.Minute * time.Duration(
			s.cfg.AuthConfig.TTLToken)).
		Unix()
}
