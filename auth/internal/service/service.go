package service

import (
	"app/access/pkg/access"
	access_pb "app/access/pkg/access/gen"
	"app/auth/internal/config"
	db "app/auth/internal/database"
	"app/auth/internal/models"
	pb "app/auth/pkg/proto/gen"
	"app/internal/token"
	"context"
	"crypto/sha1"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service - work with users and authorization
type Service interface {

	// Sign in -
	// request token by login and password.
	// Return token and user model or error
	SignIn(ctx context.Context,
		login, password string) (token string,
		user *models.User, err error)

	// List -
	// request list users by last user id
	// (Using in query WHERE id > $1)
	// and limit.
	// Return user list
	List(ctx context.Context,
		lastID, limit uint64) (users []*models.User, err error)

	// User -
	// request user by id.
	// Return user model
	User(ctx context.Context,
		userID uint64) (user *models.User, err error)

	// Create -
	// request to create new user with login and password.
	// User name = login.
	// Return error or nil
	Create(ctx context.Context, login, password string) error

	// Delete -
	// request to delete user by id.
	// Return error or nil
	Delete(ctx context.Context, userID uint64) error

	// Update stream -
	// Usage for sending create/update/delete user.
	UpdateStream(ctx context.Context,
		c chan *pb.UpdateStreamResponse) error

	// Change user image -
	// request to change user image by id from claims.
	// Return error or nil
	ChangeUserImage(ctx context.Context,
		image string) error

	// Change user -
	// request to change user name/login/email by id from claims.
	// Return error or nil
	ChangeUser(ctx context.Context,
		name, login, email string) error

	// Change user password -
	// request to change user password by id from claims.
	// Return error or nil
	ChangeUserPassword(ctx context.Context,
		oldPassword, newPassword string) error
}

type Auth struct {
	db     db.Repository
	cfg    *config.Config
	m      *token.Source
	access *access.Client
}

func New(db db.Repository, cfg *config.Config, m *token.Source, access *access.Client) *Auth {
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
	return fmt.Sprintf("%x", hash.Sum([]byte(s.getAdminSalt())))
}

func (s *Auth) getExp() int64 {
	return time.Now().Add(time.Minute * time.Duration(s.cfg.AuthConfig.TTLToken)).Unix()
}

func (s *Auth) Token(ctx context.Context, login, password string) (string, *models.User, map[string][]string, error) {

	log.Printf("login: %s, password: %s", login, password)

	// Admin auth
	if login == s.cfg.AuthConfig.AdminName && password == s.cfg.AuthConfig.AdminPassword {
		log.Println("is admin")
		claims := &token.Claims{
			UserId:  0,
			Login:   "",
			Exp:     time.Now().Add(time.Minute * time.Duration(s.cfg.AuthConfig.TTLToken)).Unix(),
			IsAdmin: true,
		}

		token, err := s.m.Sign(claims)
		if err != nil {
			return "", nil, nil, err
		}

		return token, &models.User{ID: 0, Name: "admin", Login: "admin"}, nil, nil
	}

	log.Println("is user")

	passSHA1 := models.PasswordSHA1(password, s.getAdminSalt())
	user, err := s.db.Users().SignIn(ctx, login, passSHA1)
	if err != nil {
		return "", nil, nil, status.Error(codes.InvalidArgument, "invalid user credentials")
	}

	accessServices, err := s.access.UserAccess(&access_pb.UserAccessRequest{
		Id: user.ID,
	})
	if err != nil {
		return "", nil, nil, err
	}

	services := make(map[string][]string)

	for _, s := range accessServices.Access {
		services[s.Name] = append(services[s.Name], s.Methods...)
	}

	claims := &token.Claims{
		UserId:         user.ID,
		Login:          user.Login,
		IsAdmin:        false,
		Exp:            time.Now().Add(time.Minute * time.Duration(s.cfg.AuthConfig.TTLToken)).Unix(),
		ServicesAccess: services,
	}

	token, err := s.m.Sign(claims)
	if err != nil {
		return "", nil, nil, err
	}

	return token, user, services, nil
}
