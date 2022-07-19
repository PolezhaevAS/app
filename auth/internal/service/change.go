package service

import (
	"context"

	"app/internal/token"

	"app/auth/internal/models"
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
	oldPass, newPass string, isReset bool, userID uint64) (err error) {
	oldPassword := s.getPasswordSHA1(oldPass)
	newPassword := s.getPasswordSHA1(newPass)
	if isReset {
		err = s.db.ChangePassword(ctx, userID,
			oldPassword, newPassword, isReset)
		if err != nil {
			return
		}
		return
	}

	claims, _ := token.ClaimsFromCtx(ctx)
	err = s.db.ChangePassword(ctx, claims.UserID,
		oldPassword, newPassword, isReset)
	if err != nil {
		return
	}

	return
}
