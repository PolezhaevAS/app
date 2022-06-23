package service

import (
	"app/access/pkg/access"
	access_pb "app/access/pkg/access/gen"
	"app/auth/internal/config"
	db "app/auth/internal/database"
	"app/auth/internal/models"
	"app/internal/token"
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	Token(ctx context.Context, login, password string) (string, *models.User, error)
	List(ctx context.Context) ([]*models.User, error)
	User(ctx context.Context, id uint64) (*models.User, error)
	Create(ctx context.Context, name, login, password string) error
	Update(ctx context.Context, u *models.User) error
	Delete(ctx context.Context, id uint64) error
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

func (s *Auth) Token(ctx context.Context, login, password string) (string, *models.User, error) {

	passSHA1 := models.PasswordSHA1(password, s.getAdminSalt())
	user, err := s.db.Users().SignIn(ctx, login, passSHA1)
	if err != nil {
		return "", nil, status.Error(codes.InvalidArgument, "invalid user credentials")
	}

	accessServices, err := s.access.UserAccess(&access_pb.UserAccessRequest{
		Id: user.ID,
	})
	if err != nil {
		return "", nil, err
	}

	services := make(map[string][]string)

	for _, s := range accessServices.Access {
		services[s.Name] = append(services[s.Name], s.Methods...)
	}

	claims := &token.Claims{
		UserId:         user.ID,
		Login:          user.Login,
		Exp:            time.Now().Add(time.Minute * time.Duration(s.cfg.AuthConfig.TTLToken)).Unix(),
		ServicesAccess: services,
	}

	token, err := s.m.Sign(claims)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

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
