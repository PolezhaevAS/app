package service

import (
	"context"

	"app/internal/service"
	"app/internal/token"

	access_pb "app/access/pkg/access/gen"

	"app/auth/internal/models"
)

func (s *Auth) SignIn(ctx context.Context,
	login, password string) (token string,
	user models.User, access map[string][]string, err error) {
	if s.cfg.AuthConfig.AdminName == login &&
		s.cfg.AuthConfig.AdminPassword == password {
		return s.signInAdmin(login)
	}

	return s.signInUser(ctx, login, password)
}

func (s *Auth) signInAdmin(login string) (tokenString string,
	user models.User, access map[string][]string, err error) {
	access = make(map[string][]string)
	services := service.Services()

	for _, service := range services {
		var methods []string
		for _, method := range service.Methods {
			methods = append(methods, method.MethodName)
		}
		access[service.ServiceName] = methods
	}

	claims := &token.Claims{
		UserID:  0,
		IsAdmin: true,
		Exp:     s.getExp(),
		Access:  access,
	}

	tokenString, err = s.m.Sign(claims)
	if err != nil {
		return
	}

	return tokenString, models.User{
		ID:    0,
		Name:  login,
		Login: login,
		Email: "",
	}, access, nil
}

func (s *Auth) signInUser(ctx context.Context,
	login, password string) (tokenString string,
	user models.User, access map[string][]string, err error) {
	access = make(map[string][]string)
	pass := s.getPasswordSHA1(password)
	user, err = s.db.SignIn(ctx, login, pass)
	if err != nil {
		return
	}

	answer, err := s.access.UserAccess(
		&access_pb.UserAccessRequest{Id: user.ID})
	if err != nil {
		return
	}

	for service, methods := range answer.Access {
		access[service] = methods.Name
	}

	claims := &token.Claims{
		UserID:  user.ID,
		IsAdmin: false,
		Exp:     s.getExp(),
		Access:  access,
	}

	tokenString, err = s.m.Sign(claims)
	if err != nil {
		return
	}

	return
}
