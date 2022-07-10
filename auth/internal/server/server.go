package server

import (
	"app/auth/internal/service"
	pb "app/auth/pkg/proto/gen"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	s service.Service
	pb.UnimplementedAuthServer
}

func New(s service.Service) *Server {
	return &Server{
		s: s,
	}
}

func (s *Server) SignIn(ctx context.Context,
	req *pb.SignInRequest) (*pb.SignInResponse, error) {
	token, user, access, err := s.s.
		SignIn(ctx, req.GetLogin(), req.GetPassword())
	if err != nil {
		return &pb.SignInResponse{}, status.Error(codes.Aborted, err.Error())
	}

	var userAccess map[string]*pb.Methods
	for service, methods := range access {
		userAccess[service] = &pb.Methods{Name: methods}
	}

	return &pb.SignInResponse{
		Token:  token,
		User:   user.Proto(),
		Access: userAccess,
	}, nil
}
