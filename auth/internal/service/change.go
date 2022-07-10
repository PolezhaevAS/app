package service

import (
	"app/auth/internal/models"
	"app/internal/token"
	"context"
)

func (s *Auth) ChangeUser(ctx context.Context,
	name, login, email string) (err error) {
	claims, _ := token.ClaimsFromCtx(ctx)
	user := models.User{
		ID:    claims.UserID,
		Name:  name,
		Login: login,
		Email: email,
	}

	err = s.db.Update(ctx, user)
	if err != nil {
		return
	}

	return
}

func (s *Auth) ChangeUserPassword(ctx context.Context,
	oldPass, newPass string) (err error) {
	claims, _ := token.ClaimsFromCtx(ctx)
	err = s.db.ChangePassword(ctx, claims.UserID, oldPass, newPass)
	if err != nil {
		return
	}

	return
}
