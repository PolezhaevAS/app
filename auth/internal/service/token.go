package service

import (
	"app/auth/internal/models"
	"app/internal/token"
	"context"
	"log"
)

func (s *Auth) SignIn(ctx context.Context,
	login, password string) (string, *models.User, error) {
	// Check if it's admin user
	if login == s.cfg.AuthConfig.AdminName &&
		password == s.cfg.AuthConfig.AdminPassword {
		log.Println("admin user")
	}
	log.Println("user")

	passwordSHA1 := s.getPasswordSHA1(password)
	user, err := s.db.Users().SignIn(ctx, login, passwordSHA1)
	if err != nil {
		return "", nil, err
	}

	// todo: need return map[ServiceName][]MethodsName

	// accessServices, err := s.access.UserAccess(
	// 	&access_pb.UserAccessRequest{Id: user.ID})
	// if err != nil {
	// 	return "", nil, err
	// }

	access := make(map[string][]string)

	claims := &token.Claims{
		UserID:  user.ID,
		IsAdmin: false,
		Exp:     s.getExp(),
		Access:  access,
	}

	token, err := s.m.Sign(claims)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (s *Auth) admin() (string, error) {
	access := make(map[string][]string)
	claims := &token.Claims{
		UserID:  0,
		IsAdmin: true,
		Exp:     s.getExp(),
		Access:  access,
	}

	return s.m.Sign(claims)
}

func (s *Auth) user(ctx context.Context, login, pass string) (string, error) {
	password := s.getPasswordSHA1(pass)
	user, err := s.db.Users().SignIn(ctx, login, password)
	if err != nil {
		return "", err
	}

	access := make(map[string][]string)

	claims := &token.Claims{
		UserID:  user.ID,
		IsAdmin: false,
		Exp:     s.getExp(),
		Access:  access,
	}

	return s.m.Sign(claims)
}
