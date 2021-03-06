package auth

import (
	"context"
	"errors"
	"strings"

	grpcmid "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"

	"app/internal/service"
	"app/internal/token"
)

const Scheme = "bearer"

var (
	ErrUnsupportedMethod = errors.New("unsupported method")
	ErrAccessDenied      = errors.New("access id denied")
)

type Auth struct {
	m *token.Source
	s *service.Service
}

func New(m *token.Source, s *service.Service) *Auth {
	return &Auth{
		m: m,
		s: s,
	}
}

func (a *Auth) getClaims(ctx context.Context) (*token.Claims, error) {
	token, err := grpcauth.AuthFromMD(ctx, Scheme)
	if err != nil {
		return nil, err
	}

	claims, err := a.m.Parse(token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func methodName(m string) string {
	return m[strings.LastIndex(m, "/")+1:]
}

// UnaryServerInterceptor for gRPC server.
func (a *Auth) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
		resp interface{}, err error) {

		var (
			method = methodName(info.FullMethod)
			open   = a.s.OpenApi(method)
		)

		if open {
			return handler(ctx, req)
		}

		var claims *token.Claims
		if claims, err = a.getClaims(ctx); err != nil {
			return
		}

		if claims.IsAdmin {
			return handler(ctx, req)
		}

		if err = claims.CheckAccess(a.s.Name(), method); err != nil {
			return
		}

		return handler(token.CtxWithClaims(ctx, claims), req)
	}
}

// StreamServerInterceptor for gRPC server.
func (a *Auth) StreamServerInterceptor() grpc.StreamServerInterceptor {

	return func(srv interface{}, stream grpc.ServerStream,
		info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {

		var (
			method = methodName(info.FullMethod)
			open   = a.s.OpenApi(method)
		)

		if open {
			return handler(srv, stream)
		}

		var ctx = stream.Context()

		var claims *token.Claims
		if claims, err = a.getClaims(ctx); err != nil {
			return
		}

		if claims.IsAdmin {
			return handler(srv, stream)
		}

		if err = claims.CheckAccess(a.s.Name(), method); err != nil {
			return
		}

		var wrapped = grpcmid.WrapServerStream(stream)
		wrapped.WrappedContext = token.CtxWithClaims(ctx, claims)

		return handler(srv, wrapped)
	}
}
