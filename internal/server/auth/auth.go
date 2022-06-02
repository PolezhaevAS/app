package auth

import (
	"app/internal/service"
	"app/internal/token"
	"context"
	"errors"
	"strings"

	grpcmid "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
)

const Scheme = "bearer"

var (
	ErrUnsupportedMethod = errors.New("unsupported method")
	ErrAccessDenied      = errors.New("access id denied")
)

type Auth struct {
	m           *token.Source
	descService *service.Service
}

func New(m *token.Source, service *service.Service) *Auth {
	return &Auth{
		m:           m,
		descService: service,
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

func (a *Auth) checkAccess(claims *token.Claims, method string) error {
	var userMethods, ok = claims.ServicesAccess[a.descService.Name]

	if !ok {
		return ErrAccessDenied
	}

	for _, userMethod := range userMethods {
		if method == userMethod {
			return nil
		}
	}

	return ErrAccessDenied
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
			method         = methodName(info.FullMethod)
			access, exists = a.descService.Methods[method]
		)

		// This need to check access
		if !exists {
			err = ErrUnsupportedMethod
			return
		}

		if !access {
			return handler(ctx, req)
		}

		var claims *token.Claims
		if claims, err = a.getClaims(ctx); err != nil {
			return
		}

		if err = a.checkAccess(claims, method); err != nil {
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
			method         = methodName(info.FullMethod)
			access, exists = a.descService.Methods[method]
		)

		// This need to check access
		if !exists {
			err = ErrUnsupportedMethod
			return
		}

		if !access {
			return handler(srv, stream)
		}

		var ctx = stream.Context()

		var claims *token.Claims
		if claims, err = a.getClaims(ctx); err != nil {
			return
		}

		if err = a.checkAccess(claims, method); err != nil {
			return
		}

		var wrapped = grpcmid.WrapServerStream(stream)
		wrapped.WrappedContext = token.CtxWithClaims(ctx, claims)

		return handler(srv, wrapped)
	}
}
