package token

import (
	"context"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Services access - map[service][]methods
type Claims struct {
	UserId         uint64
	Login          string
	ServicesAccess map[string][]string
	Exp            int64
}

type claimsContextKey struct{}

func CtxWithClaims(ctx context.Context, c *Claims) context.Context {
	return context.WithValue(ctx, claimsContextKey{}, c)
}

func ClaimsFromCtx(ctx context.Context) (*Claims, bool) {
	claims, b := ctx.Value(claimsContextKey{}).(*Claims)
	return claims, b
}

func (c *Claims) Valid(h *jwt.ValidationHelper) error {
	if h == nil {
		h = jwt.DefaultValidationHelper
	}

	exp := jwt.NewTime(float64(c.Exp))
	if err := h.ValidateExpiresAt(exp); err != nil {
		return status.Error(codes.Unauthenticated, fmt.Sprintf("authorization token expired %v", err.Error()))
	}

	if c.UserId <= 0 {
		return status.Error(codes.Unauthenticated, "invalid authorization token: missing user ID in the token")
	}

	return nil
}

func (c *Claims) Access(service, method string) error {
	var methods, ok = c.ServicesAccess[service]

	if !ok {
		return errors.New("no access to service")
	}

	for _, claimMethod := range methods {
		if method == claimMethod {
			return nil
		}
	}

	return errors.New("no access")

}
